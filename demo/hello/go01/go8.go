package main

import "fmt"

func main8() {
	// defer函数延迟执行，等到函数返回了再执行
	defer fmt.Println("world")
	fmt.Println("hello")

	//------------------defer栈：先进后出--------------------

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	/*
		要从栈的角度理解，碰到defer关键字之后就压栈，函数执行完毕就开始弹栈
			运行结果：
		hello
		counting
		done
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0
		world
	*/
}

/*
defer
defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
*/
