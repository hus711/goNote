//专门处理对表的增删改查操作
package main

import (
	"encoding/json"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
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
		t, _ := os.Create(POSTER_DIR + "/" + movieMsg.Title + ".jpg")
		defer t.Close()
		f, err := http.Get(movieMsg.Poster)
		check(err)
		// defer f.Body.Close()
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
	fileInfoArr, err := ioutil.ReadDir(POSTER_DIR)
	check(err)
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		if ext := path.Ext(fileInfo.Name()); ext != ".jpg" {
			continue
		}
		images = append(images, fileInfo.Name())
	}

	router.GET("/show", func(c *gin.Context) {
		renderHtml(c.Writer, "show", images)
	})
	router.GET("/", func(c *gin.Context) {
		renderHtml(c.Writer, "list", images)
	})

	router.GET("/view", func(c *gin.Context) {
		imageId := c.Query("id") //获取图片id
		//拼接图片相对路径
		imagePath := POSTER_DIR + "/" + imageId
		fmt.Println("图片路径：" + imagePath)
		if ok := isExists(imagePath); !ok {
			c.String(http.StatusNotFound, "？ 海报未找到！", imageId)
		} else {
			//设置回复头信息
			c.Header("Content-Type", "image")
			data, err := ioutil.ReadFile(imagePath)
			check(err)
			//向服务器发送图片数据
			c.Data(http.StatusOK, "image", data)
		}
	})
	router.Run(":8080")

	fmt.Println("Web服务已经开启：8080端口、.../show  .../")
	return 0
}

//保存查询到的信息到数据中去
func SaveMsg2Db(args []string) int {
	if movieMsg.Title == "" || movieMsg.ImdbID == "" {
		fmt.Println("请先搜索得到电影数据（mn-常见电影名称）")
		return 0
	} else {
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
	save				//保存搜索到的信息到数据库
	mn					//常见电影名称提示
	q					//关闭程序
	h					//帮助信息
`)
	return 0
}

func Quit(args []string) int {
	fmt.Println("退出程序")
	CloseDb()
	return 1
}

//检查是否出错，出错就panic
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//专门用来渲染函数模板
func renderHtml(w io.Writer, tmpl string, locals []string) {
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
