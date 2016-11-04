/*
指针
Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。

var p *int
& 符号会生成一个指向其作用对象的指针。

i := 42
p = &i
* 符号表示指针指向的底层的值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“非直接引用”。

与 C 不同，Go 没有指针运算。
*/

package main

import "fmt"

func main1() {
	i, j := 42, 2701

	p := &i         // point to i =====>将i的地址赋值给p,也就是p指向了i
	fmt.Println(*p) // read i through the pointer   ====>*p是指读取指针p地址指向的位置存储的值
	*p = 21         // set i through the pointer   =====》操作的是i的值
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j   ===》j的地址赋值给p
	*p = *p / 37   // divide j through the pointer   ==>将p指向的值除以37然后再赋值给p指向的位置
	fmt.Println(j) // see the new value of j
}
