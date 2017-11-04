package main

import (
	"flag"
	"fmt"
)

//命令行标志是指定命令行程序选项的常用方法, 例如，在wc -l中，-l是命令行标志
//Go提供了一个支持基本命令行标志解析的标志包
func main() {
	wordPtr := flag.String("word", "foo", "a string") //关键词   默认值  描述

	numbPtr := flag.Int("numb", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

//命令行执行：go run test4.go  -world=opt -numb=7 -fork -svar=flag
/*
命令行提示：
flag provided but not defined: -word
Usage of C:\Users\ADMINI~1\AppData\Local\Temp\go-build276010638\command-line-arguments\_obj\exe\test4.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -world string
        a string (default "foo")
exit status 2
*/
