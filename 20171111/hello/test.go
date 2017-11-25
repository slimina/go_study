package hello

import (
	"errors"
)

/*
Go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC。
但Go的RPC包是独一无二的RPC，它和传统的RPC系统不同，它只支持Go开发的服务器与客户端之间的交互，
因为在内部，它们采用了Gob来编码。
Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
函数必须是导出的(首字母大写)
必须有两个导出类型的参数，
第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
函数还要有一个返回值error
正确的RPC函数格式如下：
func (t *T) MethodName(argType T1, replyType *T2) error
T、T1和T2类型必须能被encoding/gob包编解码。
任何的RPC都需要通过网络来传递数据，Go RPC可以利用HTTP和TCP来传递数据
*/
type MunUtil int8

type Request struct {
	A, B int
}

type Response struct {
	Quo, Rem int
}

func (m *MunUtil) Multiply(request Request, reply *int) error {
	*reply = request.A * request.B
	return nil
}
func (m *MunUtil) Divide(request Request, response *Response) error {
	if request.B == 0 {
		return errors.New("divide by zero")
	}
	response.Quo = request.A / request.B
	response.Rem = request.A % request.B
	return nil
}
