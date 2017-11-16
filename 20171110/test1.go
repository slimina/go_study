package main

import (
	"fmt"
	"strconv"
)

//字符串转化的函数strconv
func main() {

	str := make([]byte, 0, 100)
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str = strconv.AppendInt(str, 4567, 10) //整数  十进制
	fmt.Println(string(str))
	fmt.Println(string(strconv.AppendQuoteRune(str, '单')))

	//Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	//Parse 系列函数把字符串转换为其他类型
	fmt.Println(strconv.ParseInt("13", 10, 64))
	fmt.Println(strconv.ParseBool("true"))
}
