package main

import (
	"fmt"
)

//Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。
//当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
//所有的内置类型和自定义类型都是可以作为匿名字段的
type Human struct {
	name   string
	age    int
	weigth int
}

type skills []string

type Student struct {
	Human      //匿名字段，那么默认Student就包含了Human的所有字段
	skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
	name       string //最外层的优先访问，也就是当你通过student.name访问的时候，是访问student里面的字段，而不是human里面的字段。
}

func main() {

	s1 := Student{Human: Human{"Lucy", 23, 48}, speciality: "Math"}
	fmt.Println(s1)
	s1.name = "Tom"
	s1.weigth += 10
	fmt.Println(s1)

	s1.Human.age = 78
	fmt.Println(s1)

	s1.skills = []string{"C#", "Java"}
	s1.int = 112
	fmt.Println(s1)

	s1.skills = append(s1.skills, "GO", "PHP")
	fmt.Println(s1)

	s1.name = "TOM"
	fmt.Println(s1)
	fmt.Println(s1.name, s1.Human.name)
}
