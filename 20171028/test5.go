package main

import (
	"fmt"
)

//常量的值在编译的已经确定
//常量的右边必须是常量或常量表达式
//常量的表达式中函数必须是内置函数

const a int = 1
const b = 'B'
const (
	text   = "123"
	length = len(text)
	num    = b * 20
)

const c, d = 23, "hello"
const (
	e, f, h = "11", len(e), 45 * 5
)

const (
	x = 123
	y        //在定义常量组时，如果不提供初始值，则将使用上一行的表达式
	z = iota //iota是常量计数器，从0开始，组内每定义一个常量自动增加1,此时z位于第3元素，故为2
	w
) //每遇到一个const关键词，iota就会重置为0

//枚举定义，使用iota实现枚举
const (
	Red = iota //第一个不可以省略表达式
	Yellow
	Green
	Blue
)

func main() {
	fmt.Println(a, b, c, e, f, h)
	fmt.Println(x, y, z, w)
	fmt.Println(Red, Yellow, Green, Blue)
}
