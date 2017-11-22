package main

import (
	"fmt"
	"html/template"
	"os"
)

//Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象
//如果要访问当前对象的字段通过{{.FieldName}},
//但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的),否则在渲染的时候就会报错

type Friend struct {
	Name string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	t := template.New("tpl-name")
	//如果模板中输出{{.}}，这个一般应用于字符串对象，默认会调用fmt包输出字符串的内容。
	t, _ = t.Parse("hello {{.UserName}}!") //template.ParseFiles
	p := Person{UserName: "tom"}
	t.Execute(os.Stdout, p)
	fmt.Println()
	fmt.Println("-----------------------------------")
	//输出嵌套字段内容
	//用{{with …}}…{{end}}和{{range …}}{{end}}来进行数据的输出
	t = template.New("example1") //template不能重复使用
	t, _ = t.Parse(`
       hello {{.UserName}} !
       {{range .Emails}}
       		an email {{.}}
       {{end}}
       {{with .Friends}}
       {{range .}}
       		my friend name is {{.Name}}
       {{end}}
       {{end}}
		`)
	p1 := Person{UserName: "JACK", Emails: []string{"123@122.com", "344433@163.com"},
		Friends: []*Friend{&Friend{Name: "111"}, &Friend{Name: "222"}}}
	t.Execute(os.Stdout, p1)
}
