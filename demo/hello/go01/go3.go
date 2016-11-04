package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//go语言的基本数据类型
func main3() {
	const f = "%T(%v)\n" //类型（值）换行
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

/*
bool   布尔类型

string 字符串类型

int  int8  int16  int32  int64  有符号整型
uint uint8 uint16 uint32 uint64 uintptr  无符号整型

byte // uint8 的别名

rune // int32 的别名
     // 代表一个Unicode码

float32 float64  浮点型

complex64 complex128
*/
