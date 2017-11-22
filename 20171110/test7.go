package main

import (
	"fmt"
	"text/template"
)

//模板包里面有一个函数Must，它的作用是检测模板是否正确，
//例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
func main() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}
