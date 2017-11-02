package main

import (
	"fmt"
)

//之前都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢
//Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，
//当多个channel都准备好的时候，select是随机的选择一个执行的。

func test(c, quit chan int) { //多个channel
	x, y := 1, 1
	for {
		select {
		case c <- x: //未准备好，堵塞,执行default
			fmt.Println(x, y)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() { //匿名函数
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		quit <- 0 //退出
	}()

	test(c, quit)
}

//在select里面还有default语法，select其实就是类似switch的功能，
//default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
/*
select {
case i := <-c:
	// use i
default:
	// 当c阻塞的时候执行这里
}
*/
