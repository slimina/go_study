package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//grep和sed是常用的行过滤器。
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error:", err)
		os.Exit(1)
	}
	//通过控制台输入 go run test2.go
}
