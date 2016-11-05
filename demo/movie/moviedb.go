package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//初始化MySQL数据库，返回连接
func initMysql() *sql.DB {
	//打开数据库连接Open(驱动名,连接字符串)
	db, err := sql.Open("mysql", "root:chinawlx@/movie")
	if err != nil {
		fmt.Println("数据库连接失败.....")
	}
	return db
}

//向数据库中插入数据
func insertData(msg MSG) {
	db := initMysql()
	//插入数据
	stmt, err := db.Prepare("insert info set imid=?,title=?,year=?,runtime=?,actors=?,plot=?,language=?,country=?,poster=?,type=?")

	if err != nil {
		fmt.Println("插入失败。。。。。", err.Error())
		return
	}
	res, err := stmt.Exec(msg.ImdbID, msg.Title, msg.Year, msg.Runtime, msg.Actors,
		msg.Plot, msg.Language, msg.Country, msg.Poster,
		msg.Type)
	if err != nil {
		fmt.Println("插入失败。。。。。", err.Error())
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("插入失败。。。。。", err.Error())
		return
	}
	fmt.Println("插入成功 id=", id)
}

//根据新信息的id查询是否存在
func findMsg(id string) bool {
	db := initMysql()
	rows, err := db.Query("SELECT title FROM info where imid=\"" + id + "\"")
	if err != nil {
		fmt.Println("查找失败。。。", err.Error())
		return false
	}
	for rows.Next() {
		var title string
		err = rows.Scan(&title) //一定要与查询到的数据一一对应，否则会失败，最好的做法就是需要什么数据，sql语句查询的时候就查那些
		if err != nil {
			fmt.Println("电影数据已经存在,imid=" + id)
		} else {
			fmt.Println(title + "的电影数据已经保存在数据库中了")
		}
		return true
	}

	return false
}

//更新数据
func updateMsg(msg MSG) {
	flag := findMsg(msg.ImdbID)
	if !flag {
		fmt.Println(msg.Title + "的数据不存在")
		return
	}

	db := initMysql()
	stmt, err := db.Prepare("update info set title=?,year=?,runtime=?,actors=?,plot=?,language=?,country=?,poster=?,type=? where imid=?")
	if err != nil {
		fmt.Println("更新失败，", err.Error())
		return
	}
	res, err := stmt.Exec(msg.Title, msg.Year, msg.Runtime, msg.Actors,
		msg.Plot, msg.Language, msg.Country, msg.Poster,
		msg.Type, msg.ImdbID)
	if err != nil {
		fmt.Println("更新失败，", err.Error())
		return
	}
	num, err := res.RowsAffected()
	if err != nil {
		return
	} else {
		fmt.Println("更新成功，affect=", num)
	}
}

//删除信息
func deleteMsg(id string) {
	flag := findMsg(id)
	if !flag {
		fmt.Println(id + "的数据不存在")
		return
	}

	db := initMysql()
	stmt, err := db.Prepare("delete from info where imid=?")
	if err != nil {
		fmt.Println("删除失败，", err.Error())
		return
	}
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("删除失败，", err.Error())
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else {
		fmt.Println("删除成功，affect=", num)
	}
}
