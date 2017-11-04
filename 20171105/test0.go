package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//读取文件需要检查大多数调用错误
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//读文件
	dat, err := ioutil.ReadFile("./file")
	check(err)
	fmt.Println(string(dat))

	//使用 os.Open打开一个文件获取一个 os.File 值开始
	f, err := os.Open("./file")
	check(err)

	//从文件开始位置读取一些字节 ，最多读取20 字节
	b1 := make([]byte, 20)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println(n1, string(b1))

	// offset: 需要移动偏移的字节数
	//whence:给offset参数一个定义，表示要从哪个位置开始偏移；0代表从文件开头开始算起，
	//       1代表从当前位置开始算起，2代表从文件末尾算起。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2) //返回读的字节数
	check(err)
	fmt.Println(n2, o2, string(b2))

	//io
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	check(err)
	fmt.Println(n3, o3, string(b3))

	//没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = f.Seek(0, 0)
	check(err)
	//bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println(string(b4))
	//任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）。
	f.Close()

}
