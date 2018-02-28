package main

import (
	"fmt"
)

//slice常见操作 https://purewhite.io/2017/04/30/go-slice-tricks/
func main() {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int
	//另一种初始化方式
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println(numbers, numbers["ten"])
	/*
		1.map 是无序的，每次打印都会不一样，不能使用index获取，只能使用key
		2.长度不固定，和slice一样，是引用类型
		3.内置的len使用map，返回key的个数
		4.不是线程安全，在多个go-routine存取时，必须使用mutex lock机制
	*/
	var a = map[string]float64{"math": 80.9, "chinese": 70.233}
	fmt.Println(len(a))
	//map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	score, ok := a["math"]
	fmt.Println(score, ok)
	score1, ok1 := a["xxx"]
	fmt.Println(score1, ok1)
	//删除key为C的元素
	delete(a, "chinese")
	fmt.Println(a)
	//引用类型
	b := make(map[string]int)
	b["a"] = 1
	b["b"] = 2
	fmt.Println(b)
	c := b
	c["c"] = 3
	fmt.Println(b)
	fmt.Println(c)
	//遍历map,for配合range可以用于读取slice和map的数据
	for k, v := range c {
		fmt.Println(k, "=>", v)
	}
	// Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错,
	//在这种情况下, 可以使用_来丢弃不需要的返回值
	for _, v := range c {
		fmt.Println("=>", v)
	}
	// make用于内建类型（map、slice 和channel）的内存分配
	//new用于各种类型的内存分配

	//make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T
	//make初始化了内部的数据结构，填充适当的值。
	//make返回初始化后的（非零）值。

	//new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
	//new返回指针

	var p = new(int)
	fmt.Println(p)
	q := new(map[string]int)
	fmt.Println(q)
	var m = make(map[string]int)
	fmt.Println(m)
}
