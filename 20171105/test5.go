package main

import (
	"fmt"
	"os"
	"strings"
)

//环境变量是一种用于将配置信息传递到Unix程序的通用机制。
func main() {

	os.Setenv("test", "123")

	fmt.Println(os.Getenv("test"))
	fmt.Println()
	for _, e := range os.Environ() {
		arr := strings.Split(e, "=")
		fmt.Println(arr[0], "=>", arr[1])
	}
}
