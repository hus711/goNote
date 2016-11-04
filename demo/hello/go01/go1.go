package main

import (
	"fmt"
	"math"
)

//函数
func main1() {
	fmt.Println(math.Pi)
	fmt.Println(add2(3, 4))

	//参数的赋值一样可以多个一起
	a, b := swap("hello", "go")
	fmt.Println(a, b)

	fmt.Println(split(1))
}

//函数的参数类型，变量名称 变量类型。返回值在后面
func add1(x int, y int) int {
	return x + y
}

//连续类型是同一个，可以直接省略
func add2(x, y int) int {
	return x + y
}

//返回值可以有多个
func swap(x, y string) (string, string) {
	return y, x
}

//函数的返回值可以被命名，并且当做变量来使用，当返回空的时候也就是直接返回
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
