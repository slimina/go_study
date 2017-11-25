package main

import (
	hello "./hello"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

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
		go rpc.ServeConn(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
