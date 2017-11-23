package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

/*
在Go语言的net包中有一个类型TCPConn，这个类型可以用来作为客户端和服务器端交互的通道，
他有两个主要的函数：
func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error)
TCPConn可以用在客户端和服务器端来读写数据。

还有我们需要知道一个TCPAddr类型，他表示一个TCP的地址信息，他的定义如下：
type TCPAddr struct {
	IP IP
	Port int
}
在Go语言中通过ResolveTCPAddr获取一个TCPAddr
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)

Go语言中通过net包中的DialTCP函数来建立一个TCP连接，并返回一个TCPConn类型的对象
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
laddr表示本机地址，一般设置为nil
raddr表示远程的服务地址

*/

func main() {
	fmt.Println(os.Args)
	ip := net.ParseIP("127.0.0.1")
	fmt.Println(ip) //type IP []byte

	//net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only),TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个).
	//addr表示域名或者IP地址
	ipadd, err := net.ResolveTCPAddr("tcp4", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipadd)
	//模拟http

	tcpAdd, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:12345")
	checkError(err)
	//获取连接
	conn, err := net.DialTCP("tcp", nil, tcpAdd)
	checkError(err)
	defer conn.Close()
	bb, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	fmt.Println(bb)
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	fmt.Println(result)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//控制TCP连接
//TCP有很多连接控制函数，我们平常用到比较多的有如下几个函数：
//func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)
//设置建立连接的超时时间，客户端和服务器端都适用，当超过设置时间时，连接自动关闭。
//func (c *TCPConn) SetReadDeadline(t time.Time) error
//func (c *TCPConn) SetWriteDeadline(t time.Time) error
//用来设置写入/读取一个连接的超时时间。当超过设置时间时，连接自动关闭。
//func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
//设置keepAlive属性，是操作系统层在tcp上没有数据和ACK的时候，会间隔性的发送keepalive包，
//操作系统可以通过该包来判断一个tcp连接是否已经断开，在windows上默认2个小时没有收到数据
//和keepalive包的时候人为tcp连接已经断开，
