package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

//初始化操作，将模板缓存起来
func init() {
	//读取模板文件夹
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		//模板名
		templateName = fileInfo.Name()
		//返回文件扩展名（包含点），不是以.html结尾的不加入模板中
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

//检查是否出错，出错就panic
func check(err error) {
	if err != nil {
		panic(err)
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

//上传图片
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//get请求，直接调用上传的网页
		renderHtml(w, "upload", nil)
	}
	//浏览器里面调用的 post请求
	if r.Method == "POST" {
		//拿到图片文件数据
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		//创建文件,返回文件对象，可以用于IO
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		//复制文件数据到文件流中
		_, err = io.Copy(t, f)
		check(err)
		//重定向
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

//查看某个图片
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//从请求中拿到id对应的信息
	imageId := r.FormValue("id")
	//拼接图片相对路径
	imagePath := UPLOAD_DIR + "/" + imageId
	if ok := isExists(imagePath); !ok {
		http.NotFound(w, r)
		return
	}

	//设置回复头信息
	w.Header().Set("Content-Type", "image")
	//向服务器发送图片数据
	http.ServeFile(w, r, imagePath)
}

//列出所有的上传图片
func listHandler(w http.ResponseWriter, r *http.Request) {
	//遍历上传文件夹
	fileInfoArr, err := ioutil.ReadDir("./uploads")
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
	fileInfoArr, err := ioutil.ReadDir("./uploads")
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
				log.Println(string(debug.Stack()))
			}
		}()

		//执行业务逻辑
		fn(w, r)
	}
}

//参数：http请求的多路转接器（将URL与注册模式列表进行匹配），模式，静态路径，
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			fi, err := os.Stat(file)
			if err != nil || fi.IsDir() {
				http.NotFound(w, r)
				return
			}
		}
		//想浏览器传送文件
		http.ServeFile(w, r, file)
	})
}

func main() {
	//创建一个多路转接器
	mux := http.NewServeMux()

	//css js文件暂时不导入
	//staticDirHandler(mux, "/assets/", "./public", 0)

	//注册不同的路径
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/show", safeHandler(showPictureHandler))
	//监听本地8080端口
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
