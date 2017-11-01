package main

import (
	"fmt"
)

//接口继承
type Print interface {
	fmt.Stringer
}

type Human struct {
	name string
}

func (h Human) String() string {
	return h.name
}

func main() {
	var out = Human{"TOM"}
	fmt.Println(out.String())

	var a Print
	a = out
	fmt.Println(a.String())

	if val, ok := a.(Print); ok {
		fmt.Println("out => Print", ":", val)
	}

	if val, ok := a.(fmt.Stringer); ok { //接口父类
		fmt.Println("out => fmt.Stringer", ":", val)
	}
}
