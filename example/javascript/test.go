package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

/*
otto 是一个 golang 写的 JavaScript 解释器，
可以用于在 golang 中执行 JavaScript，并且可以让“虚拟机”中的 JS 调用 golang 函数，
实现了服务端 golang 和 JavaScript 的互操作。
*/
func main() {
	vm := otto.New()
	vm.Set("x", 123)
	vm.Run(`
		abc = 2+x;
		console.log(abc);
	`)
	fmt.Println(vm.Get("abc"))
	value, err := vm.Get("abc")
	if err != nil {
		fmt.Println(err)
	}
	val, _ := value.ToInteger()
	fmt.Println(val)

	vm.Set("str", "abc123")
	fmt.Println(vm.Run("str.length"))

	//设置函数
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return otto.Value{}
	})
	//带有返回值的函数
	vm.Set("sayHello2", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})

	vm.Run(`
		sayHello("xxxyyy");
		sayHello();
		result = sayHello2(2);
		`)
	fmt.Println(vm.Get("result"))

	//编译
	script, err := vm.Compile("", `var abc; if (!abc) abc = 0; abc += 2; abc;`)
	vm.Run(script)
	a, _ := vm.Get("abc")
	fmt.Println(a.ToInteger())

}
