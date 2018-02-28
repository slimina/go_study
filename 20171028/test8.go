package main

import (
	"fmt"
	"runtime"
)

//只需设置 GOOS 和 GOARCH 两个环境变量就能生成所需平台的Go程序。
// 编译它： $ GOOS=darwin GOARCH=386 go build test.go  生成运行在OS X上的程序。
func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}

/*
可用的OS和ARCH的值如下：
$GOOS	$GOARCH
darwin	386
darwin	amd64
darwin	arm
darwin	arm64
dragonfly	amd64
freebsd	386
freebsd	amd64
freebsd	arm
linux	386
linux	amd64
linux	arm
linux	arm64
linux	ppc64
linux	ppc64le
netbsd	386
netbsd	amd64
netbsd	arm
openbsd	386
openbsd	amd64
openbsd	arm
plan9	386
plan9	amd64
solaris	amd64
windows	386
windows	amd64
不同的操作系统下的库可能有不同的实现， 比如syscall库。go build没有内置的#define或者预处理器之类的处理平台相关的代码取舍， 而是采用tag和文件后缀的方式实现。
在文件的头部增加tag:
// +build darwin freebsd netbsd openbsd
可以有多个tag,之间是AND的关系
// +build linux darwin
// +build 386
注意tag和package中间需要有空行分隔
// +build !linux

package mypkg
.........
文件后缀方式
以_$GOOS.go为后缀的文件只在此平台上编译，其它平台上编译时就当此文件不存在。完整的后缀如：
_$GOOS_$GOARCH.go
如syscall_linux_amd64.go,syscall_windows_386.go,syscall_windows.go等。
*/
