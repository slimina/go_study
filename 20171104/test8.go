package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseFloat("12.891", 64))
	fmt.Println(strconv.ParseInt("123", 0, 32))
	fmt.Println(strconv.ParseInt("0x1c8", 0, 32))
	fmt.Println(strconv.Atoi("123")) //字符串转数字
	fmt.Println(strconv.Itoa(122))   //数字转是字符串

	_, err := strconv.Atoi("abc")
	fmt.Println(err)
}
