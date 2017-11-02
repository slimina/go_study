package main

import (
	"fmt"
)

//上一节，我们需要读取两次c，这样不是很方便
//我们可以通过range，像操作slice或者map一样操作缓存类型的channel

func test(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //close关闭channel，关闭channel之后就无法再发送任何数据了
}
func main() {

	c := make(chan int, 10)
	go test(cap(c), c)
	//for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。
	for i := range c {
		fmt.Println(i)
	}
	//测试channel是否被关闭
	v, ok := <-c
	fmt.Println(v, ok) //false 关闭
}

//记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
//记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，
//或者你想显式的结束range循环之类的
