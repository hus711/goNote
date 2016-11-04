package main

import (
	"fmt"
	"math"
)

//结构体
type Vertex struct {
	X, Y float64
}

//结构体方法，这里使用的是指针引用，可以改变值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main3() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(*v, v.Abs())
}
