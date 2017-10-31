package main

import (
	"fmt"
)

//defer 延时语句 当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。例如资源释放
//多次调用defer 采用后进先出模式
func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
