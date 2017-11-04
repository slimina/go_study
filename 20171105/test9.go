package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

//golang.org/x/xxx 这类包托管在 github.com/golang
//下载$GOPATH/src目录后,本地安装go install github.com/golang/text
//在线安装：go get -u github.com/golang/text  安装在$GOPATH/src下 需要修改目录为golang.org/x/text
func main() {
	s := []byte("你好")
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(string(d))
	}
}
