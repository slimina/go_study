package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	//==新建文件
	//func Create(name string) (file *File, err Error)
	//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。
	//func NewFile(fd uintptr, name string) *File
	//根据文件描述符创建相应的文件，返回一个文件对象
	file, _ := os.Create("test.txt")
	fmt.Println(file.Name())
	file.Close()
	os.Remove("test.txt")

	stdin := os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	defer stdin.Close()
	//==打开文件
	//func Open(name string) (file *File, err Error)
	//该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
	//func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
	//打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
	file, _ = os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	//==写文件
	//func (file *File) Write(b []byte) (n int, err Error) 写入byte类型的信息到文件
	//func (file *File) WriteAt(b []byte, off int64) (n int, err Error) 在指定位置开始写入byte类型的信息
	//func (file *File) WriteString(s string) (ret int, err Error)  写入string信息到文件
	file.WriteString("abc123")
	file.Close()

	//== 读文件
	//func (file *File) Read(b []byte) (n int, err Error) 读取数据到b中
	//func (file *File) ReadAt(b []byte, off int64) (n int, err Error) 从off开始读取数据到b中

	file, _ = os.Open("test.txt")
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

/*
打开标记：
O_RDONLY：只读模式(read-only)
O_WRONLY：只写模式(write-only)
O_RDWR：读写模式(read-write)
O_APPEND：追加模式(append)
O_CREATE：文件不存在就创建(create a new file if none exists.)
O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
O_TRUNC：打开并清空文件
文件权限（unix权限位）：只有在创建文件时才需要，不需要创建文件可以设置为 0。os库虽然提供常量，但是我一般直接写数字，如0664。
如果你需要设置多个打开标记和unix权限位，需要使用位操作符"|"

Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
*/
