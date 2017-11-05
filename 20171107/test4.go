package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	client.Password = "abc123"
	//client.Auth("abc123") //密码验证未实现
	client.Db = 0
	//字符串操作
	err := client.Set("test", []byte("123"))
	val, err := client.Get("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(val))
	client.Del("test")

	vals := []string{"a", "b", "c", "d", "e", "f"}
	for _, v := range vals {
		client.Rpush("list", []byte(v))
	}

	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del("l")
}
