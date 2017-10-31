package main

import (
	"fmt"
	"os"
)

//go 不能抛出异常，而是使用了panic和recover机制
/*
-- 类似抛出异常
Panic是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。
当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。
在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，
此时程序退出。恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。

--类似捕获异常
Recover是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。
在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入恐慌，
调用recover可以捕获到panic的输入值，并且恢复正常的执行。
*/

func getSystemProperty(key string) {
	var val = os.Getenv(key) //获取环境变量
	if val == "" {
		panic("环境变量不存在")
	}
}

//捕获异常
func throwsPanic(f func(string), key string) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			b = true //返回值
		}
	}()
	f(key) //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

func main() {
	fmt.Println(throwsPanic(getSystemProperty, "user"))
}
