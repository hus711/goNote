package main

import (
	"fmt"
)

//var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面
var c, python, java bool

//初始化变量
var m, n int = 1, 2

//变量学习
func main2() {
	var i int
	fmt.Println(i, c, python, java)

	//直接初始化可以省略掉类型
	var c, python, java = true, false, "go"
	fmt.Println(c, python, java)

	// 短声明变量,:=可以替代var或者用于类型明确的地方,这种声明不能用在函数外
	k := 3
	fmt.Println(k)

}
