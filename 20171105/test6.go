package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//当需要一个运行的Go进程可访问的外部进程时,可使用Go的经典exec函数来实现。
//需要一个我们想要执行的二进制的绝对路径，所以将使用exec.LookPath来找到它(可能是/bin/ls)
//此程序仅限在 Linux 类系统上执行
func main() {
	binary, err := exec.LookPath("java")
	if err != nil {
		panic(err)
	}
	fmt.Println(binary)

	args := []string{"java", "-version"}
	env := os.Environ()

	err1 := syscall.Exec(binary, args, env)

	if err1 != nil {
		panic(err1)
	}
}
