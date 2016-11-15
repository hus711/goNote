package main

// import (
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/go-sql-driver/mysql"
// )

// //MySQL数据库测试
// func main() {
// 	db, err := sql.Open("mysql", "root:chinawlx@/movie")
// 	if err != nil {
// 		fmt.Println("数据库连接失败.....")
// 	}

// 	//插入数据
// 	stmt, err := db.Prepare("insert info set imid=?,title=?,year=?,runtime=?,actors=?,plot=?,language=?,country=?,poster=?,type=?,sqlstr=?")

// 	if err != nil {
// 		fmt.Println("插入数据失败.....")
// 	}
// 	_, err = stmt.Exec("tt0111161", "The Shawshank Redemption", "1994", "142 min", "Tim Robbins, Morgan Freeman, Bob Gunton, William Sadler",
// 		"ds",
// 		"English", "USA", "https://images-na.ssl-images-amazon.com/images/M/MV5BODU4MjU4NjIwNl5BMl5BanBnXkFtZTgwMDU2MjEyMDE@._V1_SX300.jpg",
// 		"movie", "sql0")
// 	if err != nil {
// 		fmt.Println("插入失败。。。。。", err.Error())
// 	}

// 	// id, err := res.LastInsertId()
// 	// if err != nil {
// 	// 	fmt.Println("插入失败。。。。。")
// 	// }
// 	// fmt.Println(id)

// }

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
