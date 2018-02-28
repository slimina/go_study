package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//func TempDir(dir, prefix string) (name string, err error)
//TempDir会在你给出的dir中创建一个新的以prefix为前缀的名字随机且不重复的临时目录
//，如果dir是空，则使用系统默认的临时目录（比如在mac上是/var/，在linux上是/tmp/），并且返回这个目录的绝对路径（完整路径）。
//func TempFile(dir, prefix string) (f *os.File, err error)
//只不过返回的是一个文件指针而已。不过要记得，如果你需要文件的后缀名有一定要求的话，需要自己去调用os.Rename函数进行修改。

//调用这两个函数一定要记得自己去删除，不然go不会主动帮你删除掉你创建出来的临时目录和文件
func main() {
	file, err := ioutil.TempFile("", "aaa")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())
}
