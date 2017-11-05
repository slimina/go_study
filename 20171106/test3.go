package main

import (
	"fmt"
	"html/template"
	"os"
	tpl "text/template"
)

func main() {

	//html/template里面带有下面几个函数可以帮你转义
	/*
		func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
		func HTMLEscapeString(s string) string //转义s之后返回结果字符串
		func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
	*/
	b := []byte("<script>alert(123)</script>")
	template.HTMLEscape(os.Stdout, b)
	fmt.Println()
	fmt.Println(template.HTMLEscapeString("<script>alert(123)</script>"))
	//html/template包默认帮你过滤了html标签，但有时候不想过滤，请使用text/template。

	t, err := tpl.New("test").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		fmt.Println(err)
	}
	//或者使用template.HTML类型
	fmt.Println()
	t, err = tpl.New("test").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//转换成template.HTML后，变量的内容也不会被转义
	err = t.ExecuteTemplate(os.Stdout, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	if err != nil {
		fmt.Println(err)
	}
}
