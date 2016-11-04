/*
Go 语言切片(Slice)
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
*/
package main

import "fmt"

func main5() {
	p := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("=====================")
	//对slice进行切片
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

	fmt.Println("=====================")
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
	fmt.Println("=====================")

	//追加元素

	printSlice("追加元素前", p)
	p = append(p, 10, 11)
	printSlice("追加元素", p)
}

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

/*
	p == [0 1 2 3 4 5 6 7 8 9]
	p[0] == 0
	p[1] == 1
	p[2] == 2
	p[3] == 3
	p[4] == 4
	p[5] == 5
	p[6] == 6
	p[7] == 7
	p[8] == 8
	p[9] == 9
	=====================
	p[1:4] == [1 2 3]
	p[:3] == [0 1 2]
	p[4:] == [4 5 6 7 8 9]
	=====================
	a len=5 cap=5 [0 0 0 0 0]
	b len=0 cap=5 []
	c len=2 cap=5 [0 0]
	d len=3 cap=3 [0 0 0]
	=====================
	追加元素前 len=10 cap=10 [0 1 2 3 4 5 6 7 8 9]
	追加元素 len=12 cap=20 [0 1 2 3 4 5 6 7 8 9 10 11]
*/
