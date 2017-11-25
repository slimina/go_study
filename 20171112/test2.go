package main

import (
	"errors"
	"fmt"
	"math"
)

/*
error类型是一个接口类型
type error interface {
	Error() string
}
error是一个内置的接口类型,很多内部包里面用到的 error是errors包下面的实现的私有结构errorString
type errorString struct {
	s string
}
func (e *errorString) Error() string {
	return e.s
}
通过errors.New把一个字符串转化为errorString，以得到一个满足接口error的对象
func New(text string) error {
	return &errorString{text}
}
*/

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(f), nil
}
func main() {
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
}
