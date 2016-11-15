package main

import (
	"bufio"
	"fmt"
	"github.com/flosch/pongo2"
	"io/ioutil"
	"log"
	//"net/http"
	"os"
	"path"
	"strings"
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
	//Response string `gorm:"-"` //回复是否成功 true,同时数据库忽略掉这个字段
	ImdbID string `gorm:"column:imid"` //信息的唯一id，作为数据库的操作条件
}

var movieMsg MSG //查询到的消息主体

var templates = make(map[string]*pongo2.Template) //所有加载到的模板文件

func init() {
	//读取模板数据
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
		t := pongo2.Must(pongo2.FromFile(templatePath))
		templates[templateName] = t
	}
}

/*
	1.搜索电影-->显示电影信息 or fail msg
	2.命令操作拿到的信息（保存、下载海报、数据库CRUD、海报的查看与删除、开启和关闭web服务）
	3.开启了web服务之后，就可以通过浏览器网页端来访问搜索到的数据
*/
func main() {
	fmt.Println("go语言电影搜索Demo：gorm、gin、pongo2综合使用")
	Help(nil)

	r := bufio.NewReader(os.Stdin)

	handlers := GetCommandHandlers()

	for { // 循环读取用户输入
		fmt.Print("Command> ")
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command:", tokens[0])
		}
	}

}
