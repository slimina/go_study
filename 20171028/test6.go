package main

import (
	"fmt"
)

/*
运算符从左到右
优先级 从高到底
^  !

* / % 	<<   >>  &  &^

+ - | ^
== !=  <  <=  >= >
<-  用于channel
&&
||

指针：不支持运算符和->，直接使用"."选择符操作指针目标对象成员
使用操作符“&“ 取变量地址，使用“*”通过指针间接访问对象
默认值为nil，而非null

递增减语句
++ 和 -- 作为语句而非表达式
====
6 : 0110
11: 1011
--------
&   0010  //有2个都为1就为1
|   1111  //有一个为1就为1
^   1101  //相同为0 不同为1
&^  0100  //第二个数为1是 值为0  为0时值为1
*/

//枚举表示计算机存储单位
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(1 << 10 >> 10)
	fmt.Println(6 &^ 11)
	fmt.Println(1 ^ 2)

	fmt.Println(B, KB, MB, GB, TB, PB, EB, ZB, YB)

	var a = 123
	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p)
	var b = *&a
	fmt.Println(b)

	b++
	// 不能写b=b++，++ 和 -- 作为语句而非表达式单独存在
	fmt.Println(b)

}
