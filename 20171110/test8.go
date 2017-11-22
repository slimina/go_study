package main

import (
	"fmt"
	"os"
	"text/template"
)

/*
Go语言中通过如下的语法来申明
{{define "子模板名称"}}内容{{end}}
通过如下方式来调用：
{{template "子模板名称"}}
*/
func main() {
	t, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	t.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	t.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	t.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	t.Execute(os.Stdout, nil)
}

//同一个集合类的模板是互相知晓的，如果同一模板被多个集合使用，则它需要在多个集合中分别解析
