package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
Match模式只能用来对字符串的判断
regexp包中含有三个函数用来判断是否匹配，如果匹配返回true，否则返回false
func Match(pattern string, b []byte) (matched bool, error error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, error error)
func MatchString(pattern string, s string) (matched bool, error error)
*/
func main() {
	m, _ := regexp.Match("[0-9]+", []byte("21234"))
	fmt.Println(m)

	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	m, _ = regexp.MatchReader("^[a-zA-Z]+$", r)
	fmt.Println(m)

	m, _ = regexp.MatchString("^\\d+", "32435443")
	fmt.Println(m)
}
