package main

import (
	. "fmt"
	"time"
)

//有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时
func main() {

	c := make(chan int)
	b := make(chan bool)
	go func() {
		for { //死循环
			select {
			case v := <-c:
				Println(v)
			case t := <-time.After(5 * time.Second): //sleep 5s After返回 chan time.Time
				Println("timeout", t)
				b <- true
				break
			}

		}
	}()
	Println(time.Now())
	a := <-b
	Println(a)
}
