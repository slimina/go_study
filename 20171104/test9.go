package main

import (
	"fmt"
	"net/url"
)

//URL提供了一种统一的方法来定位资源。
//包括方案，身份验证信息，主机，端口，路径，查询参数和查询片段。

func main() {
	s := "ftp://username:passwd@www.baidu.com:80/query?key=hello"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Host, u.Hostname(), u.Port())
	fmt.Println(u.Path, u.Fragment)
	fmt.Println(u.RequestURI(), u.Fragment, u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m, m["key"])
}
