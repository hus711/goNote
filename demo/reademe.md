1. hello—学习go指南的时候写的例子
   1. 直接运行即可
2. cgss-Go语言编程第四章的例子
   1. 功能：棋牌服务游戏
   2. 运行：go run cgss.go（或者.cgss），其它看命令
3. photoweb-Go语言编程第五章的例子
   1. 功能：上传图片、打开图片
   2. 运行方式：终端  go run photoweb.go即可(或者./photoweb)
   3. 浏览器访问：[查看所有名称列表](http://localhost:8080/) [上传](http://localhost:8080/upload)  [图片展示](http://localhost:8080/upload)
4. movie-Go语言圣经中的例子，改造了下
   1. 功能：搜索电影名字，显示电影的相关数据，并且下载海报。最后将拿到的数据存到MySQL数据库中去，学习数据库的增删改查操作。之后将拿到的数据，展示到网页上去，形成一个网页版本的搜索功能。
   2. 运行：go run movie.go 3+idiots(没有后面的参数默认下载一个)