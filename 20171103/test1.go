package main

import (
	"errors"
	"fmt"
)

//错误是最后一个返回值，并有类型：error
//errors.New使用给定的错误消息构造基本错误值。错误位置中的nil值表示没有错误。
func f1(n int) (int, error) {
	if n == 0 {
		return -1, errors.New("n is 0")
	}
	return n * n, nil
}

//error是一个内置接口。通过对它们实现Error()方法，可以使用自定义类型作为错误
//在这种情况下，使用＆argError语法构建一个新的结构，为两个字段arg和prob提供值。
type argError struct {
	arg  int
	prob string
}

//实现error的Error接口
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(n int) (int, error) {
	if n == 0 {
		return -1, &argError{n, "n is 0"}
	}
	return n * n, nil
}

func main() {
	n, e := f1(0)
	fmt.Println(n, e)
	a, b := f2(0)
	fmt.Println(a, b)
}
