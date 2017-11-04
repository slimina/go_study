package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

type Person struct {
	name string
	age  int
}

func main() {
	p(s.Contains("abcdef", "bc"))
	p(s.Count("test", "t"))
	p(s.HasPrefix("test", "te"))
	p(s.HasSuffix("test", "st"))
	p(s.Index("test", "es"))
	p(s.Join([]string{"a", "b"}, "-"))
	p(s.Repeat("ab", 5))
	p(s.Replace("foo", "o", "0", -1), s.Replace("foo", "o", "0", 1)) //替换1个
	p(s.Split("a:b:c", ":"))
	p(s.ToLower("TEst"), s.ToUpper("TEst"))
	p()
	p(len("abc"), "abc"[1])

	person := Person{"Jack", 23}
	p(person)
	fmt.Printf("%v\n", person)
	fmt.Printf("%+v\n", person) //结构体的字段名
	fmt.Printf("%#v\n", person) //源代码片段
	fmt.Printf("%T\n", person)  //类型
}
