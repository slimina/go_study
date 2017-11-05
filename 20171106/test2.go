package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, "=>", strings.Join(v, ","))
	}

	fmt.Fprintf(w, "hello xxxx")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])     //返回的是数组
		fmt.Println("password:", r.Form.Get("password")) //返回单值，因为如果字段不存在，通过该方式获取的是空值
	}

	//表单验证
	//1.汉字 regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname"))
}
func main() {

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe err :", err)
	}
}
