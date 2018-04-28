package main

import (
	"fmt"
	"net/http"
	"time"
)

//https://my.oschina.net/ureyishere/blog/967907

func main() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
		Timeout:       time.Second * 5, //从连接建立一直到获得请求结果
	}

	resp, err := client.Get("https://www.google.com/")

	fmt.Println(resp)

}
