package main

import (
	"fmt"
)

var ( //不能省略var
	a1     = 123
	a2     = "abc"
	a3, a4 = true, 321212 //多变量赋值
	// 不可以使用 a := 123
)

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	var a = 123
	var b string = "hello"
	fmt.Println(a)
	fmt.Println(b)
	a = 234
	fmt.Println(a)
	//简单声明赋值变量,不能用于全局变量  :代替var
	d := true
	fmt.Println(d)

	var b1, b2, b3 int   //多变量声明
	b1, b2, b3 = 1, 2, 3 //多变量赋值
	fmt.Println(b1, b2, b3)
	//多变量声明赋值,可以省略类型由系统判断
	var c1, c2, c3 int = 4, 5, 6
	fmt.Println(c1, c2, c3)
	d1, d2, d3 := 7, 8, 9
	fmt.Println(d1, d2, d3)

	var e, f = 123, "hello"
	fmt.Println(e, f)

	h, _, j := 1, 2, 3 //使用下划线可以作为占位符
	fmt.Println(h, j)
}
