package main

import (
	"bytes"
	"fmt"
	"github.com/robertkrimen/otto"
	"os"
)

func main() {

	f, _ := os.Open("test.js")
	defer f.Close()

	buff := bytes.NewBuffer(nil)
	//读文件
	buff.ReadFrom(f)

	runtime := otto.New()
	runtime.Run(buff.String()) //运行脚本

	a := 1
	b := 2

	jsa, err := runtime.ToValue(a)
	jsb, err := runtime.ToValue(b)
	result, err := runtime.Call("add", nil, jsa, jsb)
	if err != nil {
		panic(err)
	}
	out, err := result.ToInteger()
	fmt.Println(out)

}
