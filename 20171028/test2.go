package main

// 类型别名
import (
	"fmt"
)

type (
	byte bool
	abc  string
	文本   string
)

func main() {
	var b byte
	var a abc
	var c 文本
	c = "文本类型"
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
}
