package main

// import "fmt"

// //操作管道c，go语言的特性就会自然阻塞，等待操作完成再释放管道
// func sum(a []int, c chan int) {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	c <- sum // 将和送入 c，管道接受值
// }

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c) //关闭一个管道，表示不再接受值的输入
// }

// func main() {
// 	a := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)     //首先创建管道
// 	go sum(a[:len(a)/2], c) //运行一个轻量级的go线程,7+2+8=17
// 	go sum(a[len(a)/2:], c) // 第一个线程执行完再执行，c=17，c=-9+4+0=5
// 	x, y := <-c, <-c        // 从 c 中获取,接受管道中的值

// 	fmt.Println(x, y, x+y)

// 	//--------------------演示channel的缓冲效果

// 	//管道的容量定为10，当管道接受的值达到容量之后就阻塞，等待清空
// 	b := make(chan int, 2)
// 	b <- 1
// 	b <- 2
// 	// b  <- 3这个会阻塞，fatal error: all goroutines are asleep - deadlock!
// 	fmt.Println(<-b, <-b) //有几个值就只能发送几次，发送不存在的值就会阻塞

// 	//------------------演示range close
// 	d := make(chan int, 10)
// 	go fibonacci(cap(d), d)
// 	for i := range d { //循环取值直到没有值为止,值还能从管道中发送出来
// 		fmt.Println(i)
// 	}
// 	m, ok := <-d
// 	fmt.Println(m, ok)
// }

/*
	-5 17 12
	1 2
	0
	1
	1
	2
	3
	5
	8
	13
	21
	34
	0 false
*/

/*
channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

ch <- v    // 将 v 送入 channel ch。
v := <-ch  // 从 ch 接收，并且赋值给 v。
（“箭头”就是数据流的方向。）

和 map 与 slice 一样，channel 使用前必须创建：

ch := make(chan int)
默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步

--------------------
缓冲 channel
channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

ch := make(chan int, 100)
向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。

修改例子使得缓冲区被填满，然后看看会发生什么。
------------------------

range 和 close
发送者可以 close 一个 channel 来表示再没有值会被发送了。
接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：当没有值可以接收并且 channel 已经被关闭，那么经过

v, ok := <-ch
之后 ok 会被设置为 `false`。

循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。

注意： 只有发送者才能关闭 channel，而不是接收者。向一个已经关闭的 channel 发送数据会引起 panic。
还要注意： channel 与文件不同；通常情况下无需关闭它们。只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，例如中断一个 `range`。
*/
