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

func (h *Human) SayHello() {
	fmt.Println("hello " + h.name)
}

func (h *Human) Sing(song string) {
	fmt.Println(h.name+" sing ", song)
}

//mployee重载Human的Sayhi方法
func (e *Employee) SayHello() {
	fmt.Println("employee : hello " + e.name)
}

func (s *Student) borrowMoney() {
	fmt.Println(s.name, "borrow money")
}

func (e *Employee) SpendSalary() {

	fmt.Println(e.name, "spend salary")

}

//interface是一组method签名的组合
//interface可以被任意的对象实现。

type Men interface {
	SayHello()
	Sing(song string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney()
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary()
}

func main() {

	s := Student{Human: Human{"Tom"}, school: "school1"}
	s.Sing("love")
	s.SayHello()
	s.borrowMoney()

	var a Student
	a = s
	a.Sing("aaaa")
	a.SayHello()

	fmt.Println(&s, &a, a)
}
