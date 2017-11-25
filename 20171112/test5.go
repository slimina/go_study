package main

import (
	"fmt"
	"time"
)

//Go内部已经内置支持了GDB,目前支持调试Go程序的GDB版本必须大于7.1。
//纯go代码使用delve可以很好的进行Go代码调试:https://github.com/derekparker/delve

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}

	close(c)
}
func main() {
	bus := make(chan int)
	go counting(bus)

	for count := range bus {
		fmt.Println(count)
	}
}

//编译文件，生成可执行文件gdbfile:
//go build -gcflags "-N -l" test5.go

//通过gdb命令启动调试：
//gdb test5

//启动之后首先看看这个程序是不是可以运行起来，只要输入run命令回车后程序就开始运行
//而后参考GDB调试.md命令 调试程序,如：b 23表示在第23行设置了断点
