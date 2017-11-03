package main

import (
	"fmt"
	"time"
)

//select 超时
func main() {
	c1 := make(chan int, 1)

	go func() {
		fmt.Println("c1 start ...")
		time.Sleep(time.Second * 2)
		fmt.Println("c1 end ...")
		c1 <- 122
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1): //1秒超时
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		fmt.Println("c2 start ...")
		time.Sleep(time.Second * 2)
		fmt.Println("c2 end ...")
		c2 <- "result 2"
		c2 <- "eqqqqqqqqqq"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3): //3秒超时
		fmt.Println("timeout 2")
	}
	fmt.Println(<-c2)
}
