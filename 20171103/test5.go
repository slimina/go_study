package main

import (
	"fmt"
	"time"
)

//tickers是用于定期做一些事情。 周期性
//停止后，它不会在其通道上接收任何更多的值
func main() {
	ticker := time.NewTicker(time.Millisecond * 600)
	go func() {
		for t := range ticker.C { //周期性接收
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop() //停止
	fmt.Println("ticker is stoped")
}
