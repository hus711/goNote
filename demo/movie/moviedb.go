package main

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"strings"
)

//结果信息
type Result struct {
	Title string
	Year  string
	Imid  string
}

var db gorm.DB
var err error

//初始化数据库连接
func init() {
	db, err = gorm.Open("mysql", "root:chinawlx@/movie?charset=utf8&parseTime=True&loc=Local")
	check(err)
}

//向数据库中插入数据
func insertData() {
	//插入数据
	//stmt, err := db.Prepare("insert info set imid=?,title=?,year=?,runtime=?,actors=?,plot=?,language=?,country=?,poster=?,type=?")

	db.Table("info").Create(&movieMsg)
	// 	"imid=?,title=?,year=?,runtime=?,actors=?,plot=?,language=?,country=?,poster=?,type=?",
	// 	msg.ImdbID,msg.Title,msg.Year,msg.Runtime,msg.Actors,msg.Plot,msg.Language,msg.Country,msg.Poster,msg.Type
	// )
}

//根据新信息的id查询是否存在
func findMsg(id string) MSG {
	var msg MSG
	db.Find(&msg, "imid=?", id)
	return msg
}

//查找所有的信息
func FindAllMsg() []Result {
	msgs := make([]Result, 100)

	db.Table("info").Select("title,year,imid").Scan(&msgs)
	return msgs
}

//更新数据
func updateMsg(msg MSG) {
	//	stmt, err := db.Prepare("update info set title=?,year=?,runtime=?,actors=?,plot=?,language=?,country=?,poster=?,type=? where imid=?")
	// res, err := stmt.Exec(msg.Title, msg.Year, msg.Runtime, msg.Actors,
	// 	msg.Plot, msg.Language, msg.Country, msg.Poster,
	// 	msg.Type, msg.ImdbID)
}

//删除信息
func DeleteMsgByImid(id string) {
	db.Table("info").Delete(&Result{}, "imid = ?", id)
}

//清空info表
func DeleteAllMsg() {
	db.Table("info").Delete(&Result{})
}

//关闭数据库连接
func CloseDb() {
	db.Close()
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
