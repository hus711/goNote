package main

import "fmt"

func main6() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	for i := 0; i < 20; i++ {
		sum -= i
	}
	fmt.Println(sum)

	//变量不能重复定义
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//死循环 for{ todo }

	//if语句
	if sum > 1024 {
		fmt.Println("sum > 1024 ", "sum = ", sum)
	} else {
		fmt.Println("sum <= 1024", "sum = ", sum)
	}

	//if 语句前面可以执行一个语句,里面定义的变量是局部变量，可以在else中用
	if v := 3; v < 5 {
		fmt.Println(v)
	}

}

/*
for
Go 只有一种循环结构——`for` 循环。

基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。
*/
