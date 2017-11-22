package main

import (
	"os"
	"text/template"
)

//在Go模板里面如果需要进行条件判断，那么我们可以使用和Go语言的if-else语法类似的方式来处理
//如果pipeline为空，那么if就认为是false
//if里面无法使用条件判断 if里面只能是bool值

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	//pipelines 管道类似linux | ，Go语言模板最强大的一点就是支持pipe数据，在Go语言里面任何{{}}里面的都是pipelines数据
	//{{. | html}}  html转义

	//在模板使用过程中需要定义一些局部变量，我们可以在一些操作中申明局部变量，
	//例如with``range``if过程中申明局部变量，这个变量的作用域是{{end}}之前
	t1 := template.New("template test")
	t1 = template.Must(t1.Parse(`{{with $x := "output" | printf "%q"}}{{$x}}{{end}}`))
	t1.Execute(os.Stdout, nil)
	t1 = template.New("template test")
	t1 = template.Must(t1.Parse(`{{with $x := "output"}}{{printf "%q" $x}}{{end}}`))
	t1.Execute(os.Stdout, nil)
	t1 = template.New("template test")
	t1 = template.Must(t1.Parse(`{{with $x := "output"}}{{$x | printf "%q"}}{{end}}`))
	t1.Execute(os.Stdout, nil)
}
