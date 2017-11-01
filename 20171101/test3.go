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

type Element interface{} //自定义类型，空接口,可以存储任意对象

type List []Element //别名 slice

type Person struct {
	name string
}

func (p Person) String() string {
	return "name:" + p.name
}

func main() {

	list := make(List, 4)
	list[0] = 123
	list[1] = "abc"
	list[2] = Person{name: "TOM"}
	fmt.Println(list)
	for idx, el := range list {
		if val, ok := el.(int); ok {
			fmt.Println(idx, "=>", "int", val)
		} else if val, ok := el.(string); ok {
			fmt.Println(idx, "=>", "string", val)
		} else if val, ok := el.(Person); ok {
			fmt.Println(idx, "=>", "Person", val)
		} else {
			fmt.Println(idx, "###", el)
		}

	}
	fmt.Println("=================================================")
	//switch ：`element.(type)`语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用`comma-ok`。
	for idx, el := range list {
		switch val := el.(type) {
		case int:
			fmt.Println(idx, "=>", "int", val)
		case string:
			fmt.Println(idx, "=>", "string", val)
		case Person:
			fmt.Println(idx, "=>", "Person", val)
		default:
			fmt.Println(idx, "###", el)
		}
	}
}
