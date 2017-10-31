package main

import (
	"fmt"
)

//指针
// & 取对象地址
// * 指针对象
func main() {
	var a [2]int
	var b [4]int
	fmt.Println(a, b)
	fmt.Println(&a, &b, *&b)

	c := 45
	fmt.Println(c, &c, *&c)
}
