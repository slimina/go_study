package main

import (
	"fmt"
	"regexp"
)

/*
Compile:解析正则表达式
func Compile(expr string) (*Regexp, error)
func CompilePOSIX(expr string) (*Regexp, error)
func MustCompile(str string) *Regexp
func MustCompilePOSIX(str string) *Regexp
	CompilePOSIX和Compile的不同点在于POSIX必须使用POSIX语法，它使用最左最长方式搜索，
而Compile是采用的则只采用最左方式搜索(例如[a-z]{2,4}这样一个正则表达式，
应用于"aa09aaa88aaaa"这个文本串时，CompilePOSIX返回了aaaa，而Compile的返回的是aa)。
前缀有Must的函数表示，在解析正则语法的时候，如果匹配模式串不满足正确的语法则直接panic，
而不加Must的则只是返回错误。
*/
func main() {
	s := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")
	//查找符合正则表达式第一个
	fmt.Println(string(re.Find([]byte(s))))
	fmt.Println("----------------")
	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	for _, v := range re.FindAll([]byte(s), -1) {
		fmt.Println(string(v))
	}
	fmt.Println("----------------------")
	//查找符合条件的index位置,开始位置和结束位置
	idx := re.FindIndex([]byte(s))
	fmt.Println(idx)
	//查找符合条件的所有的index位置 ,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	arr := re.FindAllIndex([]byte(s), -1)
	fmt.Println(arr)
	fmt.Println("----------------------")
	re2, _ := regexp.Compile("am(.*)lang(.*)")
	//查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，
	//第三个是第二个()里面的
	for _, v := range re2.FindSubmatch([]byte(s)) {
		fmt.Println(string(v))
	}
	fmt.Println("----------------")
	//查询每段的开始和结束为止
	fmt.Println(re2.FindSubmatchIndex([]byte(s)))

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(s), -1)
	for _, v := range submatchall[0] { //实际只匹配到一个
		fmt.Println(string(v))
	}

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(s), -1)
	fmt.Println(submatchallindex)
}

/*
搜索的函数:
func (re *Regexp) Find(b []byte) []byte
func (re *Regexp) FindAll(b []byte, n int) [][]byte
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
func (re *Regexp) FindAllString(s string, n int) []string
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
func (re *Regexp) FindIndex(b []byte) (loc []int)
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
func (re *Regexp) FindString(s string) string
func (re *Regexp) FindStringIndex(s string) (loc []int)
func (re *Regexp) FindStringSubmatch(s string) []string
func (re *Regexp) FindStringSubmatchIndex(s string) []int
func (re *Regexp) FindSubmatch(b []byte) [][]byte
func (re *Regexp) FindSubmatchIndex(b []byte) []int
上面这18个函数我们根据输入源(byte slice、string和io.RuneReader)不同还可以继续简化成如下几个，
其他的只是输入源不一样，其他功能基本是一样的：
func (re *Regexp) Find(b []byte) []byte
func (re *Regexp) FindAll(b []byte, n int) [][]byte
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
func (re *Regexp) FindIndex(b []byte) (loc []int)
func (re *Regexp) FindSubmatch(b []byte) [][]byte
func (re *Regexp) FindSubmatchIndex(b []byte) []int
*/
