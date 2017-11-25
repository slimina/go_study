package main

import (
	hello "./hello"
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Synchronous call
	request := hello.Request{1, 2}
	var reply int
	err = client.Call("MunUtil.Multiply", request, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("request: %d*%d=%d\n", request.A, request.B, reply)

	var response hello.Response

	err = client.Call("MunUtil.Divide", request, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("--------")
	fmt.Printf("request: %d/%d=%d remainder %d\n", request.A, request.B, response.Quo, response.Rem)
}
