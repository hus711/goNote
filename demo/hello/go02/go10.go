package main

import "fmt"

//递归一定要出口
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

func main10() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
}
