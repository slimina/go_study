package main

import (
	"fmt"
)

//interface的变量里面可以存储任意类型的数值(该类型实现了interface),判断存储的对象了类型
/*
1.Comma-ok断言
Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
*/
func main() {
	//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
}
