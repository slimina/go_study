package main

import (
	"fmt"
)

//interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数(包括method)提供了一些额外的思考，
//我们是不是可以通过定义interface参数，让函数接受各种类型的参数。
/*
fmt.print接口
type Stringer interface { //接口定义
	String()
}
*/
type Human struct {
	name string
}

//通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "hello " + h.name
}

func main() {
	Bob := Human{"Bob"}
	fmt.Println("This Human is : ", Bob)
}
