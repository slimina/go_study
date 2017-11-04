package main

import (
	"fmt"
	"os"
)

//Go语言中，使用os.Exit立即退出并返回给定状态。
//当使用os.Exit时，defers不会运行。
func main() {
	defer fmt.Println("exit ...........")
	os.Exit(3)
}
