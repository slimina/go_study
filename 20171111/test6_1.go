package main

import (
	hello "./hello"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

//JSON RPC是数据编码采用了JSON，而不是gob编码
// json-rpc是基于TCP协议实现的，目前它还不支持HTTP方式。

//目前Go尚未提供对SOAP RPC的支持，但现在已经有第三方的开源实现了。
func main() {
	mu := new(hello.MunUtil)
	rpc.Register(mu)

	tcpAdd, err := net.ResolveTCPAddr("tcp", ":12345")
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAdd)
	checkErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
