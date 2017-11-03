package main

import (
	"fmt"
)

func main() {
	//if
	var a = 1
	if a > 0 {
		fmt.Println("1")
	} else if a == 0 {
		fmt.Println("0")
	} else {
		fmt.Println("-1")
	}
	var i = -1
	if i := 1; i > 0 { //i作用域只在该语句块,上一行的i被隐藏
		fmt.Println(i)
	}
	fmt.Println(i)

	//循环只有 for

	for {
		a++
		if a > 3 {
			break
		}
	}

	fmt.Println(a)

	for a < 6 {
		a++
	}
	fmt.Println(a)

	for i := 0; i < 3; i++ {
		a++
	}
	fmt.Println(a)

	//整数二进制1的个数
	i = 0
	a = 7
	for a > 0 {
		i++
		a &= (a - 1)
	}
	fmt.Println(i)

	//swicth用法,，Go里面switch默认相当于每个case最后带有break，
	//匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	//但是可以使用fallthrough强制执行后面的case代码。
	n := 1
	switch n {
	case 0:
		fmt.Println("n=0")
	case 1:
		fmt.Println("n=1")
	default:
		fmt.Println("....")
	}

	switch { //没有条件的 switch,同 switch true 一样
	case n >= 0:
		fmt.Println("n=0")
		fallthrough //继续执行下一个case
	case n >= 1:
		fmt.Println("n=1")
	}

	switch m := 1; {
	case m >= 0:
		fmt.Println("m=0")
		fallthrough //继续执行下一个case
	case m >= 1:
		fmt.Println("m=1")
	}

	//goto break continue 跳转语句
	//都可以配合标签使用，break continue配合标签跳出多层循环,标签区分大小写
	//goto是调整执行位置 建议标签放在循环后
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}
}
