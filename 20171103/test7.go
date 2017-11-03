package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter //200毫秒处理一个请求
		fmt.Println("request ", req, time.Now())
	}

	fmt.Println("===========================")
	//其他场景
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t //每200毫秒
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		burstyRequests <- i //请求
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request ", req, time.Now())
	}
}
