package main

import (
	"fmt"
)

//通道的基本发送和接收都阻塞。
//但是，可以使用select和default子句来实现非阻塞发送，接收，甚至非阻塞多路选择(select)
func main() {
	//如果消息上有可用的值，则选择将使用该值的<-message大小写。如果不是，它会立即采用默认情况。

	//对消息(message)和信号(signals)的非阻塞接收。
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("receive message : ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("send message : ", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sign := <-signals:
		fmt.Println("received signal", sign)
	default:
		fmt.Println("no activity")
	}
}
