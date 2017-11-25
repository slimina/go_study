package main

import (
	hello "./hello"
	"fmt"
	"net/http"
	"net/rpc"
)

//HTTP RPC 服务端
func main() {
	mu := new(hello.MunUtil)
	rpc.Register(mu)
	//通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
	}
}
