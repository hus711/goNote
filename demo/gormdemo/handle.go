//专门处理对表的增删改查操作
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"time"
)

var db gorm.DB
var err error

//初始化数据库连接
func init() {
	db, err = gorm.Open("mysql", "root:chinawlx@/gormdemo?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(nil)
	}
}

// 将命令和处理函数对应
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":        Help,
		"h":           Help,
		"createtable": createtable,
		"ct":          createtable,
		"deletetable": deletetable,
		"dt":          deletetable,
		"insert":      insert,
		"i":           insert,
		"query":       query,
		"q":           query,
		"qlo":         queryLO,
		"update":      update,
		"u":           update,
		"delete":      deletes,
		"d":           deletes,
		"runsql":      runsql,
		"quit":        Quit,
		"qt":          Quit,
	}
}

func Help(args []string) int {
	fmt.Println(`
Commands:
    createtable(ct) <tablename>		//创建表
    deletetable(dt) <tablename>		//删除表
    insert(i) <tablename> [<num>]	//插入数据、表名、数据量（默认1）
    query(q) <tablename> [<age>]	//查询数据、表名、年龄
    update(u) <tablename> [<age>]	//更新表、表名、年龄(默认0)
    delete（d）<tablename> [<mode><age>]	//删除数据、模式（a:删除所有.age课.、c:根据年龄）、年龄
    qlo	size page//分页查询person表：每页的数据量、第几页(0起)
    runsql <mode> <...>//运行sql语句、模式（i:增加+name+age、u:更新+id+age、q:查询+id、d:删除+id）
    quit(qt)						//退出
    help(h)							//帮助信息
`)
	fmt.Println(`
-------------------------------
	table:Person(p),Animal(a)
		`)
	return 0
}

func Quit(args []string) int {
	db.Close()
	return 1
}

//创建表
func createtable(args []string) int {
	fmt.Println("创建表：", args[1])
	switch args[1] {
	case "a":
		fallthrough
	case "Animal":
		db.CreateTable(&Animal{})
	case "p":
		fallthrough
	case "Person":
		db.CreateTable(&Person{})
	default:
		fmt.Println("表模型不存在，请重新输入或添加表模型")
	}
	return 0
}

//删除表
func deletetable(args []string) int {
	fmt.Println("删除表：", args[1])
	switch args[1] {
	case "a":
		fallthrough
	case "Animal":
		db.DropTable(&Animal{})
	case "p":
		fallthrough
	case "Person":
		db.DropTable(&Person{})
	default:
		fmt.Println("表模型不存在，请重新输入")
	}
	return 0
}

//往数据库中插入数据
func insert(args []string) int {
	fmt.Println("插入数据：", args[1])

	var num int
	if len(args) < 3 {
		num = 1
	} else {
		num, err = strconv.Atoi(args[2])
		checkErr(err)
	}

	switch args[1] {
	case "a":
		fallthrough
	case "Animal":
		for i := 0; i < num; i++ {
			age := i + 1
			animal := Animal{Birthday: time.Now(), Age: age}
			db.Create(&animal)
		}
	case "p":
		fallthrough
	case "Person":
		for i := 0; i < num; i++ {
			name := "hs" + strconv.Itoa(i)
			age := i + 1
			person := Person{Name: name, Age: age}
			db.Create(&person)
		}
	default:
		fmt.Println("表模型不存在，请重新输入")
	}
	return 0
}

//查询数据
func query(args []string) int {
	if len(args) == 1 {
		fmt.Println("退出请输入qt")
		return 0
	}

	fmt.Println("查询数据,操作表：", args[1])

	switch args[1] {
	case "a":
		fallthrough
	case "Animal":
		animals := make([]Animal, 100)
		var age string = "1"
		if len(args) == 3 {
			age = args[2]
		}
		db.Where("age_of_the_beast=?", age).Find(&animals)
		fmt.Print("查询年龄：" + age + " ")
		fmt.Println(animals)

		// anis := make([]Animal, 100)
		// db.Find(&anis)
		// fmt.Println(anis)
	case "p":
		fallthrough
	case "Person":
		person := &Person{}
		db.First(&person)
		fmt.Println("查询第一个的名字：" + person.Name)

		person1 := &Person{} //不能使用上面哪一个，不会赋值进去
		db.Last(&person1)
		fmt.Println("查询最后一个的名字：" + person1.Name)

		person2 := &Person{}
		db.Where("name = ?", "hs0").First(&person2)
		fmt.Println("条件查询第一个name=hs0的： ", person2)
	default:
		fmt.Println("表模型不存在，请重新输入")
	}
	return 0
}

//更新数据库数据
func update(args []string) int {
	fmt.Println("更新数据，操作表：", args[1])

	var age int
	if len(args) < 3 {
		age = 0
	} else {
		age, err = strconv.Atoi(args[2])
		checkErr(err)
	}

	switch args[1] {
	case "a":
		fallthrough
	case "Animal":
		//更新某个字段
		animal := &Animal{}
		db.First(&animal)
		fmt.Println(animal)
		db.Model(&animal).Where("age_of_the_beast = ?", animal.Age).Update("age_of_the_beast", age)
		animals := make([]Animal, 100)
		db.Where("age_of_the_beast = ?", age).Find(&animals)
		fmt.Println(animals)
	case "p":
		fallthrough
	case "Person":
		person := &Person{}
		//首先查询出来，填充数据
		db.First(&person)
		person.Age = age //修改年龄
		//更新数据
		db.Save(person) //Save update value in database, if the value doesn't have primary key, will insert it
		//再查找出第一个age数据，打印出来
		db.First(&person, "age=?", age)
		fmt.Println("更新后：", person)
	default:
		fmt.Println("表模型不存在，请重新输入")
	}
	return 0
}

//删除数据库数据
func deletes(args []string) int {
	fmt.Println("删除数据，操作表：", args[1])
	mode := args[2]

	var age string = "1"
	if mode == "c" {
		//指定部分删除
		age = args[3]
		checkErr(err)
	}

	switch args[1] {
	case "a":
		fallthrough
	case "Animal":
		if mode == "a" {
			animal := &Animal{}
			//主键缺省的情况下就是删除所有数据
			db.Delete(&animal)
			fmt.Println("删除所有")
		} else {
			fmt.Println("删除指定的年龄：" + age)
			animal := &Animal{}
			db.Delete(&animal, "age_of_the_beast=?", age)
		}
	case "p":
		fallthrough
	case "Person":
		if mode == "a" {
			person := &Person{}
			//主键缺省的情况下就是删除所有数据
			db.Delete(&person)
			fmt.Println("删除所有")
		} else {
			fmt.Println("删除指定的年龄：" + age)
			person := &Person{}
			db.Delete(&person, "age=?", age)
		}
	default:
		fmt.Println("表模型不存在，请重新输入")
	}
	return 0
}

//分页查询：输入三个参数：qlo、size、page
func queryLO(args []string) int {
	if len(args) != 3 {
		fmt.Println("分页查询参数异常")
		return 0
	}

	fmt.Println("分页查询数据,操作person表")
	size, _ := strconv.Atoi(args[1])
	page, _ := strconv.Atoi(args[2])
	persons := make([]Person, 100)
	db.Offset(page * size).Limit(size).Select("name, age").Find(&persons)
	fmt.Println(persons)
	return 0
}

//运行sql语句
func runsql(args []string) int {
	switch args[1] {
	case "i": //提供name 和age  runsql i lisi 99
		sql := "INSERT INTO person(`name`, `age`) VALUES (?,?);"
		db.Exec(sql, args[2], args[3])
	case "d": //提供主键即可
		sql := "delete from person where id =?"
		db.Exec(sql, args[2])
	case "u": //更新提供主键和年龄 runsql u 10 20
		sql := "UPDATE person SET age=? WHERE id =?"
		db.Exec(sql, args[3], args[2])
	case "q": //查询提供主键
		sql := "SELECT name, age FROM person"
		rows, err := db.Raw(sql).Rows()
		checkErr(err)
		defer rows.Close()

		for rows.Next() {
			var result Result
			rows.Scan(&result.Name, &result.Age)
			fmt.Println("name：" + result.Name + " 	age:" + strconv.Itoa(result.Age))
		}
	}
	return 0
}

//-----------

/*

user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

db.NewRecord(user) // => returns `true` as primary key is blank

db.Create(&user)

db.NewRecord(user) // => return `false` after `user` created
---------------------------------------------------------------
// Get first record, order by primary key
db.First(&user)
//// SELECT * FROM users ORDER BY id LIMIT 1;

// Get last record, order by primary key
db.Last(&user)
//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// Get all records
db.Find(&users)
//// SELECT * FROM users;

// Get record with primary key
db.First(&user, 10)
//// SELECT * FROM users WHERE id = 10;

// Get first matched record
db.Where("name = ?", "jinzhu").First(&user)
//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

// Get all matched records
db.Where("name = ?", "jinzhu").Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu';

db.Where("name <> ?", "jinzhu").Find(&users)

// IN
db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)

db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

----------------------
// Delete an existing record
db.Delete(&email)
//// DELETE from emails where id=10;
*/
