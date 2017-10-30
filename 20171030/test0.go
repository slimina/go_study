package main

import (
	"fmt"
)

//string 字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义
const (
	a = "hello "
	b = "world !"
)

func main() {
	fmt.Println(a + b)
	//string 是不可变的，可以通过如下方式修改

	s := "hello"
	c := []byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'a'
	s2 := string(c) // 再转换回 string 类型
	fmt.Println(s2)
	//切片
	s3 := "c" + s[1:]
	fmt.Println(s3)

	//多行字符串
	ss := `abcjkdd
		dskkkllf
		ddsdd`
	fmt.Println(ss)

}
