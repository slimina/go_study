package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

//Go语言中管理状态的主要机制是通过通道进行通信。这节使用sync/atomic包来实现由多个goroutine访问的原子计数器。
//使用一个无符号整数表示计数器(正数)

func main() {

	fmt.Println(runtime.GOMAXPROCS(10))
	//不使用原子操作
	var n uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				n++
				//time.Sleep(time.Millisecond)
				runtime.Gosched() //让出cpu
			}
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(" n= ", n)

	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddUint64(&ops, 1) //加1，原子操作
				//time.Sleep(time.Millisecond)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(ops)
	fmt.Println("ops = ", atomic.LoadUint64(&ops))
}
