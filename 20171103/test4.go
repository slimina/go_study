package main

import (
	"fmt"
	"time"
)

//Go的内置计时器和自动接收器功能
//<-timer1.C阻塞定时器的通道C，直到它发送一个指示定时器超时的值。
//如果只是想等待，可以使用time.Sleep。定时器可能起作用的一个原因是在定时器到期之前取消定时器。
func main() {
	time1 := time.NewTimer(time.Second * 2)
	fmt.Println("time1 start ...")
	<-time1.C //阻塞直到过期
	fmt.Println("Timer 1 expired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoped")
	}
}
