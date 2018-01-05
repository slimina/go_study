package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")
	tmpl.Reload(true)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	app.RegisterView(tmpl)
	app.Get("/", index)
	app.Run(iris.Addr(":8888"), iris.WithCharset("utf-8"))
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello World"})
}
