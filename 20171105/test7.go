package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//Go信号通知通过在通道上发送os.Signal值来工作
func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//signal.Notify 注册这个给定的通道用于接收特定信号。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//协程执行一个阻塞的信号接收操作
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	//程序将在这里进行等待，直到它得到了期望的信号
	fmt.Println("awaiting .......")
	<-done
	fmt.Println("exiting ......")
}

//当我们运行这个程序时，它将一直等待一个信号。
//使用 ctrl-C（终端显示为 ^C），我们可以发送一个 SIGINT 信号，这会使程序打印 interrupt 然后退出。
