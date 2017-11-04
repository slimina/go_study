package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//ioutil
	b1 := []byte("hello world")
	err := ioutil.WriteFile("data1", b1, 0600)
	check(err)

	//os
	f, err := os.Create("data2")
	check(err)
	defer f.Close()
	n1, err := f.WriteString("abcdefg---123454") //返回写入字节数
	fmt.Println(n1)

	f.Sync() //刷新到磁盘

	//bufio

	w := bufio.NewWriter(f)
	n2, err := w.WriteString("\nbuffer io")
	w.Flush()
	fmt.Println(n2)

}
