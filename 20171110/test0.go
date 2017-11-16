package main

import (
	"fmt"
	"strings"
)

//strings和strconv两个包

func main() {
	//func Contains(s, substr string) bool  字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("abcdefg", "ab"))
	fmt.Println(strings.Contains("", ""))

	//func Join(a []string, sep string) string  字符串链接，把slice a通过sep链接起来
	fmt.Println(strings.Join([]string{"hello", "world", "!"}, " "))

	//func Index(s, sep string) int 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("abcdefg", "cd"))

	//func Repeat(s string, count int) string  重复s字符串count次，最后返回重复的字符串
	fmt.Println(strings.Repeat("123", 3))

	//func Replace(s, old, new string, n int) string  在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("abcdadk", "a", "==", -1))

	//func Split(s, sep string) []string  把s字符串按照sep分割，返回slice
	fmt.Println(strings.Split("1,2,3,4,5", ","))
	fmt.Println(strings.Split(" xyz ", ""))

	//func Trim(s string, cutset string) string  在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Trim("==dsdaf==", "="))

	//func Fields(s string) []string  去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Println(strings.Fields("ss fd ss "))
	fmt.Println(strings.FieldsFunc("ss fd ss ", test))
}

func test(a rune) bool {
	fmt.Println(string(a))
	if strings.TrimSpace(string(a)) == "" {
		return true
	}
	return false
}
