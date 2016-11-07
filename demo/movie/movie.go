package main

import (
	"encoding/json"
	"fmt"
	//"io"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	//"strings"
)

const (
	BASE_URL     = "http://www.omdbapi.com" //基础数据拉取页面
	POSTER_DIR   = "./posters"              //海报的保存目录
	TEMPLATE_DIR = "./views"                //html模板的存放目录
)

//消息结构
type MSG struct {
	Title    string //标题
	Year     string //年份
	Runtime  string //时长
	Actors   string //演员
	Plot     string //简介
	Language string //语言
	Country  string //国家
	Poster   string //海报URL
	Type     string //资源类型
	Response string //回复是否成功 true
	ImdbID   string //信息的唯一id，作为数据库的操作条件
}

func (msg *MSG) string() string {
	return msg.Title
}

var movieMsg MSG
var title string = "The+Shawshank+Redemption"
var templates = make(map[string]*template.Template)

func init() {
	fileInfoDir, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName, templatePath string

	for _, fileInfo := range fileInfoDir {
		//模板名
		templateName = fileInfo.Name()
		//只需要.html结尾的文件
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		//模板路径
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		//解析模板（must解析出错直接panic）
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

//专门用来渲染函数模板
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	tmpl += ".html"
	//拿到对应的模板，然后渲染并且写到浏览器中去
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

//判断路径是否存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//列出所有的上传图片
func listHandler(w http.ResponseWriter, r *http.Request) {
	//遍历上传文件夹
	fileInfoArr, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		if ext := path.Ext(fileInfo.Name()); ext != ".jpg" {
			continue
		}
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

//显示图片
func showPictureHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		if ext := path.Ext(fileInfo.Name()); ext != ".jpg" {
			continue
		}
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "show", locals)
}

//查看某个图片,没有就不会显示（html中有view?id=...）
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//从请求中拿到id对应的信息
	imageId := r.FormValue("id")
	//拼接图片相对路径
	imagePath := POSTER_DIR + "/" + imageId
	fmt.Println("图片路径：" + imagePath)
	if ok := isExists(imagePath); !ok {
		http.NotFound(w, r)
		return
	}

	//设置回复头信息
	w.Header().Set("Content-Type", "image")
	//向服务器发送图片数据
	http.ServeFile(w, r, imagePath)
}

//使用闭包防止程序崩溃（闭包：函数和引用环境的组成的整体，函数内部返回函数）
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//延迟执行匿名函数，在每一个业务逻辑调用完毕都会调用这个匿名函数，如果出发了panic则拦截下来
		defer func() {
			//使用recover来拦截错误信息
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)

				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e.Error())

				// logging
				log.Println("WARN: panic fired in %v.panic - %v", fn, e)
			}
		}()

		//执行业务逻辑
		fn(w, r)
	}
}

//TODO:做一个循环输入的命令行程序
/*
	1.搜索电影-->显示电影信息 or fail msg
	2.命令操作拿到的信息（保存、下载海报、数据库CRUD、海报的查看与删除、开启和关闭web服务）
	3.开启了web服务之后，就可以通过浏览器网页端来访问搜索到的数据
*/
func main() {
	// fmt.Println("welcome")
	// Help(nil)

	// r := bufio.NewReader(os.Stdin)

	// handlers := GetCommandHandlers()

	// for { // 循环读取用户输入
	// 	fmt.Print("Command> ")
	// 	b, _, _ := r.ReadLine()
	// 	line := string(b)

	// 	tokens := strings.Split(line, " ")

	// 	if handler, ok := handlers[tokens[0]]; ok {
	// 		ret := handler(tokens)
	// 		if ret != 0 {
	// 			break
	// 		}
	// 	} else {
	// 		fmt.Println("Unknown command:", tokens[0])
	// 	}
	// }

	help(nil)

	//--------------------拼装URL-----------------------------
	if len(os.Args) == 2 {
		title = os.Args[1]
	}
	url := BASE_URL + "/?t=" + title + "&y=&plot=short&r=json"

	//------------------------通过搜索拿到json数据-----------------------------------
	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()
	msgByte, err := ioutil.ReadAll(res.Body)
	check(err)

	//传入byte数组，将结果存入结构体中
	json.Unmarshal(msgByte, &movieMsg)
	fmt.Println(movieMsg)

	//不存在就去下载URL中海报
	if !isExists(POSTER_DIR + "/" + movieMsg.Title + ".jpg") {

		f, err := http.Get(movieMsg.Poster)
		//去除标题中的加号
		//title = strings.Replace(title, "+", " ", -1)
		t, _ := os.Create(POSTER_DIR + "/" + movieMsg.Title + ".jpg")
		check(err)
		defer t.Close()
		b, err := ioutil.ReadAll(f.Body)
		//复制文件数据到文件流中去
		_, err = t.Write(b)
		check(err)
	} else {
		fmt.Println("海报已经存在。。。。")
	}

	//----------------------将拿到的json数据存入数据库中-----------------------------------------
	//首先查询对应的id在数据库中是否存在，不存在才插入数据
	flag := findMsg(movieMsg.ImdbID)
	if !flag {
		insertData(movieMsg)
	}
	//updateMsg(movieMsg)
	//deleteMsg(movieMsg.ImdbID)

	//-----------------------在网页上展示数据----------------------------------------------------

	//创建一个多路转接器
	mux := http.NewServeMux()
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/show", safeHandler(showPictureHandler))
	//监听本地8080端口
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

//检查是否出错，出错就panic
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func help(args []string) int {
	fmt.Println(`
	Commands:
    	sea <movie name>	//搜索电影（only English movie）
    	lip					//显示所有海报列表
    	lid					//显示所有数据库信息
    	dep <num>			//删除海报	
    	ded <imid>			//删除数据库对应的电影信息
    	deap				//删除所有海报
    	dead				//删除所有数据库数据
    	starts				//开启web服务
    	stops				//关闭web服务
    	save				//保存搜索到的信息到数据库
    	q					//关闭程序
    	h					//帮助信息
`)
	fmt.Println(`
--------------------------------------------------
	示例：
       	1.The Shawshank Redemption 肖申克的救赎
       	2.The Godfather 教父
       	3.Star Wars 星球大战
       	4.Leon-The Professional 这个杀手不太冷
       	5.Inception 盗梦空间
       	6.3 idiots 三傻大闹宝莱坞
       	7.One Flew the Cuckoo's Nest  飞越疯人院
       	8.Once Upon A Time in America 美国往事
`)

	return 0
}

/*
insert movie.info set title="hehe",year="2012",runtime="120",actors="zs"
select * from movie.info
update movie.info set title = "haha" where id="2"
delete from movie.info where id="3"
*/

/*
{
    "Title": "The Shawshank Redemption",
    "Year": "1994",
    "Rated": "R",
    "Released": "14 Oct 1994",
    "Runtime": "142 min",
    "Genre": "Crime, Drama",
    "Director": "Frank Darabont",
    "Writer": "Stephen King (short story \"Rita Hayworth and Shawshank Redemption\"), Frank Darabont (screenplay)",
    "Actors": "Tim Robbins, Morgan Freeman, Bob Gunton, William Sadler",
    "Plot": "Chronicles the experiences of a formerly successful banker as a prisoner in the gloomy jailhouse of Shawshank after being found guilty of a crime he did not commit. The film portrays the man's unique way of dealing with his new, torturous life; along the way he befriends a number of fellow prisoners, most notably a wise long-term inmate named Red.",
    "Language": "English",
    "Country": "USA",
    "Awards": "Nominated for 7 Oscars. Another 18 wins & 30 nominations.",
    "Poster": "https://images-na.ssl-images-amazon.com/images/M/MV5BODU4MjU4NjIwNl5BMl5BanBnXkFtZTgwMDU2MjEyMDE@._V1_SX300.jpg",
    "Metascore": "80",
    "imdbRating": "9.3",
    "imdbVotes": "1,711,064",
    "imdbID": "tt0111161",
    "Type": "movie",
    "Response": "True"
}
*/
