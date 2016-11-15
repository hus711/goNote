//专门处理对表的增删改查操作
package main

import (
	"encoding/json"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

var router *gin.Engine

func init() {
	//初始化一个路由引擎
	router = gin.Default()
}

//搜索电影
func SearchMovie(args []string) int {
	//--------------------拼装URL-----------------------------
	var title string
	for i, n := 1, len(args); i < n; i++ {
		if i == n-1 {
			title += args[i]
		} else {
			title += args[i] + "+"
		}
	}
	url := BASE_URL + "/?t=" + title + "&y=&plot=short&r=json"

	//------------------------通过搜索拿到json数据-----------------------------------
	fmt.Println("开始检索电影信息：" + url)
	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()
	msgByte, err := ioutil.ReadAll(res.Body)
	check(err)

	//传入byte数组，将结果存入结构体中
	json.Unmarshal(msgByte, &movieMsg)
	fmt.Println(movieMsg)

	if movieMsg.Title == "" {
		fmt.Println("电影不存在，请检查名称是否输入正确！")
		return 0
	}
	//不存在就去下载URL中海报
	if !isExists(POSTER_DIR + "/" + movieMsg.Title + ".jpg") {
		fmt.Println("下载海报中。。。")
		f, err := http.Get(movieMsg.Poster)
		check(err)
		t, _ := os.Create(POSTER_DIR + "/" + movieMsg.Title + ".jpg")
		check(err)
		defer t.Close()
		b, err := ioutil.ReadAll(f.Body)
		//复制文件数据到文件流中去
		_, err = t.Write(b)
		check(err)
		fmt.Println(movieMsg.Title + "下载成功")
	} else {
		fmt.Println(movieMsg.Title + "海报已经存在。。。。")
	}

	return 0
}

//列出已经下载的所有海报
func ListPosters(args []string) int {
	fileInfoDir, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	for _, fileInfo := range fileInfoDir {
		name := fileInfo.Name()
		//只需要.jpg结尾的文件
		if ext := path.Ext(name); ext != ".jpg" {
			continue
		}
		fmt.Println(fileInfo.Name())
	}
	return 0
}

//数据库信息列表
func ListDbMsg(args []string) int {
	msgs := FindAllMsg()
	for _, msg := range msgs {
		fmt.Println(msg)
	}
	return 0
}

//删除海报 dep 3 Idiots
func DeletePoster(args []string) int {
	var title string
	for i, n := 1, len(args); i < n; i++ {
		if i == n-1 {
			title += args[i]
		} else {
			title += args[i] + " "
		}
	}
	title += ".jpg"

	fileInfoDir, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	for _, fileInfo := range fileInfoDir {
		name := fileInfo.Name()
		//只需要.jpg结尾的文件
		if ext := path.Ext(name); ext != ".jpg" {
			continue
		}
		if name == title {
			err := os.Remove(POSTER_DIR + "/" + name)
			if err != nil {
				fmt.Println("删除失败")
			} else {
				fmt.Println("删除成功")
			}
			return 0
		}
	}
	fmt.Println("没有找到：" + title)
	return 0
}

//删除数据库中的电影信息
func DeleteDbMsg(args []string) int {
	DeleteMsgByImid(args[1])
	return 0
}

//清除所有的海报
func ClearPoster(args []string) int {
	fileInfoDir, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	for _, fileInfo := range fileInfoDir {
		name := fileInfo.Name()
		//只需要.jpg结尾的文件
		if ext := path.Ext(name); ext != ".jpg" {
			continue
		}
		err = os.Remove(POSTER_DIR + "/" + name)
		if err != nil {
			fmt.Println(name + "删除失败")
		} else {
			fmt.Println(name + "删除成功")
		}
	}

	return 0
}

//清除所有数据库信息
func ClearDbMsg(args []string) int {
	DeleteAllMsg()
	return 0
}

//开启web服务
func StartWebServer(args []string) int {
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

	// router.GET("/show", func(c *gin.Context) {
	// 	safeHandler(showPictureHandler(c.Writer, c.Request))
	// })
	// router.GET("/view", func(c *gin.Context) {
	// 	safeHandler(viewHandler(c.Writer, c.Request))
	// })
	// router.GET("/", func(c *gin.Context) {
	// 	safeHandler(listHandler(c.Writer, c.Request))
	// })

	// router.Run(":8080")

	fmt.Println("Web服务已经开启：8080端口、.../show  .../")
	return 0
}

//关闭Web服务
func StopWebServer(args []string) int {
	fmt.Println("待实现(端口无法再获取输入？)")
	return 0
}

//保存查询到的信息到数据中去
func SaveMsg2Db(args []string) int {
	if movieMsg.Title == "" {
		fmt.Println("请先搜索电影数据（mn-常见电影名称）")
		return 0
	} else {
		//fmt.Println("待实现")
		insertData()
	}

	return 0
}

// 将命令和处理函数对应
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"sea":    SearchMovie,
		"lip":    ListPosters,
		"lid":    ListDbMsg,
		"dep":    DeletePoster,
		"ded":    DeleteDbMsg,
		"cp":     ClearPoster,
		"cd":     ClearDbMsg,
		"starts": StartWebServer,
		"stops":  StopWebServer,
		"save":   SaveMsg2Db,
		"mn":     MoveNameHelp,
		"h":      Help,
		"q":      Quit,
	}
}

//常见电影名称提示
func MoveNameHelp(args []string) int {
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

func Help(args []string) int {
	fmt.Println(`
Commands:
	sea <movie name>	//联网搜索电影（only English movie）
	lip					//显示所有海报列表
	lid					//显示所有数据库信息
	dep <name>			//删除海报(省略 .jpg)	
	ded <imid>			//删除数据库对应的电影信息
	cp					//删除所有海报
	cd					//删除所有数据库数据
	starts				//开启web服务
	stops				//关闭web服务
	save				//保存搜索到的信息到数据库
	mn					//常见电影名称提示
	q					//关闭程序
	h					//帮助信息
`)
	return 0
}

func Quit(args []string) int {
	fmt.Println("退出程序")
	return 1
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

//检查是否出错，出错就panic
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//-------------------------------------------------
//专门用来渲染函数模板
func renderHtml(w http.ResponseWriter, tmpl string, locals []string) {
	tmpl += ".html"
	//拿到对应的模板，然后渲染并且写到浏览器中去
	err := templates[tmpl].ExecuteWriter(pongo2.Context{"items": locals}, w)
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
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		if ext := path.Ext(fileInfo.Name()); ext != ".jpg" {
			continue
		}
		images = append(images, fileInfo.Name())
	}
	renderHtml(w, "list", images)
}

//显示图片
func showPictureHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		if ext := path.Ext(fileInfo.Name()); ext != ".jpg" {
			continue
		}
		images = append(images, fileInfo.Name())
	}
	renderHtml(w, "show", images)
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
