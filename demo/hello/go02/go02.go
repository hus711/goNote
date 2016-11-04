/*
结构体
一个结构体（`struct`）就是一个字段的集合。

（而 type 的含义跟其字面意思相符。

-----------
结构体字段
结构体字段使用点号来访问。
------
结构体指针
结构体字段可以通过结构体指针来访问。

通过指针间接的访问是透明的
----------------
结构体文法
结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/

package main

import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

type person struct {
	name string
	age  int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

var (
	byron = person{"张三", 18}
)

func main2() {
	fmt.Println(Vertex{1, 2})

	//字段通过点来访问
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)

	//指针来操作结构体
	p := &v   //将结构体v的地址赋值给p
	p.X = 1e9 //p就可以直接调用结构体里面的字段
	p.Y = 10
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
	fmt.Println(byron)
}

/*
输出结果：
	{1 2}
	4 2
	{1000000000 10}
	{1 2} 0xc42007c020 {1 0} {0 0}
	{张三 18}
*/
