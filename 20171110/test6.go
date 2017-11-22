package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

//每一个模板函数都有一个唯一值的名字,然后与一个Go函数关联，通过如下的方式来关联
//type FuncMap map[string]interface{}

//注册模板函数 emailDeal 关联的Go函数名称是EmailDealWith
//t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})

type Friend struct {
	Name string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func main() {
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Name}}
				{{end}}
				{{end}}
				`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&Friend{Name: "111"}, &Friend{Name: "222"}}}
	t.Execute(os.Stdout, p)
}

/*
在模板包内部已经有内置的实现函数，下面代码截取自模板包里面
var builtins = FuncMap{
	"and":      and,
	"call":     call,
	"html":     HTMLEscaper,
	"index":    index,
	"js":       JSEscaper,
	"len":      length,
	"not":      not,
	"or":       or,
	"print":    fmt.Sprint,
	"printf":   fmt.Sprintf,
	"println":  fmt.Sprintln,
	"urlquery": URLQueryEscaper,
}
*/
