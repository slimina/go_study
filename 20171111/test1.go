package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

//在服务器端我们需要绑定服务到指定的非激活端口，并监听此端口，当有客户端请求到达的时候可以
//接收到来自客户端连接的请求
//func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
//func (l *TCPListener) Accept() (c Conn, err os.Error)

func main() {
	address := ":12345"
	tcpAdd, err := net.ResolveTCPAddr("tcp4", address)
	fmt.Println(tcpAdd)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAdd)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("new conn ", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
