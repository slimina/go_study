package main

import (
	"fmt"
	"os"
)

//os.Args提供对原始命令行参数的访问。请注意，此切片中的第一个值是程序的路径，os.Args [1：]保存程序的参数。
func main() {

	for _, s := range os.Args {
		fmt.Println(s)
	}
}
