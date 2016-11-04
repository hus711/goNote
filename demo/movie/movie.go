package main

import (
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	BASE_URL   = "http://www.omdbapi.com"
	POSTER_DIR = "./posters"
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
}

var movieMsg MSG
var title string = "The+Shawshank+Redemption"

//开放API接口，通过输入title将返回电影数据
func main() {
	tip()

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

	//下载URL中海报
	f, err := http.Get(movieMsg.Poster)
	t, _ := os.Create(POSTER_DIR + "/" + title + ".jpg")
	check(err)
	defer t.Close()
	b, err := ioutil.ReadAll(f.Body)
	//复制文件数据到文件流中去
	_, err = t.Write(b)
	check(err)

	//----------------------将拿到的json数据存入数据库中-----------------------------------------TODO

}

//检查是否出错，出错就panic
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func tip() {
	fmt.Println(`
	只支持搜索英文电影(空格加号代替):
    	没有参数默认值：The+Shawshank+Redemption
    	《教父》The Godfather
		《肖申克的救赎》The Shawshank Redemption
		《星球大战》Star Wars
		《Leon-The Professional 》这个杀手不太冷
		《Inception》盗梦空间
		《3 idiots》三傻大闹宝莱坞
		One Flew the Cuckoo's Nest 《飞越疯人院》
		Once Upon A Time in America《美国往事》
`)
}

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
