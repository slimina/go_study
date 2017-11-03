package main

import (
	"fmt"
)

/*
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}
1.关键字func用来声明一个函数funcName
2.函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
3.函数可以返回多个值
4.上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
5.如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
6.如果没有返回值，那么就直接省略最后的返回信息
7.如果有返回值， 那么必须在函数的外层添加return语句
8.Go语言支持匿名函数，可以形成闭包。匿名函数在想要定义函数而不必命名时非常有用。
*/

func print(a int) {
	fmt.Println(a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//多返回值
func sumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

//变参函数,变量arg是一个int的slice
func moreParams(arg ...int) {
	for _, x := range arg {
		print(x)
	}
}
func main() {
	x := 2
	y := 3
	print(max(x, y))

	n, m := sumAndProduct(x, y)
	print(n)
	print(m)

	moreParams(1, 2, 3)
	//匿名函数
	func() {
		fmt.Println(12346)
	}()

}
