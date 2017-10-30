package main

import (
	"fmt"
)

//由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度
//数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。
func main() {
	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}     //其他默认为0
	c := [...]int{1, 2, 3, 4} //自动计算长度

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	e := [...][4]int{{2, 3, 2, 4}, {12, 3}}
	fmt.Println(e)
}
