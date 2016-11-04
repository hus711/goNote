package main

// import (
// 	"fmt"
// 	"time"
// )

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	//死循环不停的判断;1.修改管道c的值，2.将quit的值输出
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit: //quit中有值就会执行
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int) //管道c
// 	quit := make(chan int)
// 	//运行线程
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c) //发送c中的值，没有阻塞
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)

// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// 	for {
// 		select {
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("BOOM!")
// 			return
// 		default:
// 			//非阻塞线程
// 			fmt.Println("    .")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// }

/*
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
	quit
	    .
	    .
	tick.
	    .
	    .
	tick.
	    .
	    .
	tick.
	    .
	    .
	tick.
	    .
	    .
	tick.
	BOOM!
*/
/*
select
select 语句使得一个 goroutine 在多个通讯操作上等待。

select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。

默认选择
当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。

为了非阻塞的发送或者接收，可使用 default 分支：

select {
case i := <-c:
    // 使用 i
default:
    // 从 c 读取会阻塞
}
*/
