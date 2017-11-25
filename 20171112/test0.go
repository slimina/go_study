package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

// go 第三方工具 https://github.com/golang
func main() {
	a := sha256.New()
	io.WriteString(a, "hello ...")
	fmt.Println(a.Sum(nil))

	b := sha1.New()
	io.WriteString(b, "abc123")
	fmt.Println(fmt.Sprintf("%x", b.Sum(nil)))

	c := md5.New()
	io.WriteString(c, "admin")
	fmt.Println(fmt.Sprintf("%x", c.Sum(nil)))

	//专家方案 scrypt,目前为止最难破解的 但计算所需时间长，而且占用的内存也多
	dk, _ := scrypt.Key([]byte("some password"), []byte("slat"), 16384, 8, 1, 32)
	fmt.Println(fmt.Sprintf("%x", dk))
}
