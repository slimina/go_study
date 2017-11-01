package main

import (
	"fmt"
)

type Human struct {
	name string
}

type Student struct {
	Human  //匿名字段Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h Human) SayHello() {
	fmt.Println("hello " + h.name)
}

func (h Human) Sing(song string) {
	fmt.Println(h.name+" sing ", song)
}

//mployee重载Human的Sayhi方法
func (e Employee) SayHello() {
	fmt.Println("employee : hello " + e.name)
}

type Men interface {
	SayHello()
	Sing(song string)
}

//任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。可以存储任意类型的数值。
//一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。
var Null interface{}

func main() {
	//定义了一个interface的变量，可以存实现这个interface的任意类型的对象
	//method Receiver 只能是值传递，不可以是引用传递
	var e = Employee{Human: Human{"Lucy"}, company: "company1"}
	e.SayHello()
	e.Sing("aaa")

	var a Men
	a = e
	a.SayHello()
	fmt.Println(a)
	var s = Student{Human: Human{"TOM"}, school: "school1"}
	a = s
	a.SayHello()
	//interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现
	x := make([]Men, 3, 6) //长度、容量
	fmt.Println(x, len(x), cap(x))
	x[0], x[1] = e, s
	for _, i := range x {
		if i != nil {
			i.SayHello()
		}
	}
	fmt.Println(x)

	fmt.Println(Null)
	c := 15
	Null = c
	fmt.Println(Null)
	Null = s
	fmt.Println(Null)
}
