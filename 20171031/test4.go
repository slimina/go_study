package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {

	var p Person
	fmt.Println(p)
	p.name = "tom"
	p.age = 30
	fmt.Println(p)

	//1.按照顺序提供初始化值
	p = Person{"Jack", 43}
	fmt.Println(p)
	//2.通过field:value的方式初始化，这样可以任意顺序
	p = Person{age: 34, name: "Lucy"}
	fmt.Println(p)
	//3.当然也可以通过new函数分配一个指针，此处P的类型为*person
	var pt = new(Person)
	fmt.Println(pt)
	pt.age = 12
	pt.name = "Tom"
	fmt.Println(pt, *pt)

	var p1 Person
	p1.age, p1.name = 34, "Join"
	p2 := Person{age: 25, name: "Bob"}

	pp, age := Older(p1, p2)
	fmt.Println(pp, age)
}

//比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age
	}
	return p2, p2.age
}
