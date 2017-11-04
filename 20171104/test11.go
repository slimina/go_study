package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := "abcsdyr!@#$%^&''=://"

	bs := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(bs)
	des, _ := base64.StdEncoding.DecodeString(bs)
	fmt.Println(string(des))

	//url base64

	bs = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(bs)
	des, _ = base64.URLEncoding.DecodeString(bs)
	fmt.Println(string(des))
}
