package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
}

func (h Human) String() string {
	return h.Name
}

func (h Human) Hello(name string) {
	fmt.Println("hello ", name)
}

//Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。我们一般用到的包是reflect包。
func main() {
	h := Human{"TOM"}
	//首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。
	t := reflect.TypeOf(h)
	fmt.Println(t) // 对象类型
	v := reflect.ValueOf(h)
	fmt.Println(v)
	fmt.Println(v.Type(), v.Kind(), v.Method(0).String()) //获取到方法  类类型v.Kind() struct
	fmt.Println(v.NumField(), v.NumMethod())              //字段数，方法数

	//转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值
	tag := t.Field(0).Type      //获取定义在struct里面的标签
	name := v.Field(0).String() //获取存储在第一个字段里面的值

	fmt.Println(tag, name)
	fmt.Println("===================")
	//修改值,需要指针

	p := reflect.ValueOf(&h)       // p为reflect.Ptr 指针类型
	fmt.Println(p.Elem().CanSet()) //是否可以修改
	q := p.Elem().Field(0)
	fmt.Println(q)
	q.SetString("====") //只能修改public的属性
	fmt.Println(v.Field(0).String())
	f := p.Elem()
	fn := f.FieldByName("Name")
	if !fn.IsValid() { //是否找到
		return
	}
	if fn.Kind() == reflect.String {
		fn.SetString("JACK")
	}
	fmt.Println(h)

	mm := f.MethodByName("Hello")
	if !mm.IsValid() {
		return
	}

	arge := []reflect.Value{reflect.ValueOf("TTTTT")}

	mm.Call(arge)
}
