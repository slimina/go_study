package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
runtime包中有几个处理goroutine的函数：
Goexit
退出当前执行的goroutine，但是defer函数还会继续调用
Gosched
让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
NumCPU
返回 CPU 核数量
NumGoroutine
返回正在执行和排队的任务总数
GOMAXPROCS
用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
*/
func main() {
	p := fmt.Println

	defer p("defer .....")
	p(runtime.NumCPU())

	p(runtime.GOMAXPROCS(1))

	go func() {
		defer p("defer func1 .....")
		p("func1 -- 1")
		runtime.Gosched()
		p("func1 -- 2")
	}()

	go func() {
		defer p("defer func2 .....")
		p("func2 -- 1")
		runtime.Goexit()
		p("func2 -- 2")
	}()
	p(runtime.NumGoroutine())
	time.Sleep(1000000000) //等待go 执行完
}
