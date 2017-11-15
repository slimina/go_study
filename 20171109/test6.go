package main

import (
	"fmt"
	"regexp"
)

//Expand 在 src 中取出相应的子串，替换掉 template 中的 $1、$2 等引用符号。
//func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
func main() {
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	// (?i) 表示所在位置右侧的表达式开启忽略大小写模式
	//  (?s) 表示所在位置右侧的表达式开启单行模式。
	// (?m) 表示所在位置右侧的表示式开启指定多行模式。更改 ^ 和 $ 的含义，以使它们分别与任何行的开头和结尾匹配，
	//       而不只是与整个字符串的开头和结尾匹配,(?m)只有在正则表达式中涉及到多行的“^”和“$”的匹配时，才使用Multiline模式
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`) //?P<arg>赋值参数
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
