package main

import (
	"fmt"
	"strconv"
)

//类型转换
func main() {
	//不存在隐式类型转转，所有类型转换必须显示声明
	//转换只能发生在2个相互兼容的类型
	var a float32 = 1.2
	b := int(a)
	fmt.Println(b)
	/*
		    不能转换
			var c bool = true
			d := int(c)
	*/

	var c = 65
	d := string(c) //作为字符编码转换
	fmt.Println(d)

	e := strconv.Itoa(c)
	fmt.Println(e)
	c, _ = strconv.Atoi("78")
	fmt.Println(c)

}
