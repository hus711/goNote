package main

import "fmt"

//切片需要一个返回值
func main6() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.空切片
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.根据需要增长
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.添加超过一个元素
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
