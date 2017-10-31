package main

import (
	"fmt"
)

//method Receiver可以是值传递，也可以是引用传递
//上一节SetColor通过指针修改color的值，如果不传Box的指针，那么SetColor接受的其实是Box的一个copy

//method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。
type Human struct {
	name string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

//在human上面定义了一个method
func (h *Human) sayHello() {
	fmt.Println("hello " + h.name)
}

//Employee 重写sayHello
func (e *Employee) sayHello() {
	fmt.Println("employee :" + e.name)
}

func main() {
	s := Student{Human{"Jack"}, "school1"}
	var e = Employee{Human{"Tom"}, "company1"}

	s.sayHello()
	e.sayHello()
}
