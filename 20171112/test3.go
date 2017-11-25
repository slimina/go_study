package main

import (
	"fmt"
	"strconv"
)

//error是一个interface,通过定义实现此接口的结构，我们就可以实现自己的错误定义
type RequestError struct {
	Code    int
	message string
}

//实现error接口
func (e RequestError) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.message
}

func CheckNil(v interface{}) error {
	if v == nil {
		return RequestError{500, "value is nill"}
	}
	return nil
}

func main() {
	err := CheckNil(nil)
	if err != nil {
		fmt.Println(err)
	}

	err, ok := err.(RequestError) //判断异常类型 例如 net.Error定义
	fmt.Println(err)
	fmt.Println(ok)
}
