//GORM的使用demo
package main

//导入GORM所要使用的包：http://jinzhu.me/gorm/database.html#connecting-to-a-database
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("欢迎进入GORM学习DEMO")
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
