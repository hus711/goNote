package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main5() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)

	//-----------类型推导
	v := 42
	fmt.Printf("v is of type %T\n", v)

	//---------常量，不能使用:=语法
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//------------------数值常量：由上下文确定类型

	fmt.Println("=================")
	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big)) //数据越界
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*
类型转换
表达式 T(v) 将值 v 转换为类型 `T`。
与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换,不加报错
*/
