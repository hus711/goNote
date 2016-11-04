/*
	方法
		Go 没有类。然而，仍然可以在结构体类型上定义方法。
		方法接收者 出现在 func 关键字和方法名之间的参数中。
		--------

方法（续）
你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。

但是，不能对来自其他包的类型或基础类型定义方法。
*/

package main

import (
	"fmt"
	"math"
)

//结构体
type Vertex struct {
	X, Y float64
}

//浮点数类型
type MyFloat float64

//改类型对应的方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//结构体方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main1() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
