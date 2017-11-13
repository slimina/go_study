package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Go的JSON包中有如下函数
//func Unmarshal(data []byte, v interface{}) error  解析json
//func Marshal(v interface{}) ([]byte, error)  生成json

//能够被赋值的字段必须是可导出字段(即首字母大写）。
//同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略

/*
JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组
Go类型和JSON类型的对应关系如下：
bool 代表 JSON booleans,
float64 代表 JSON numbers,
string 代表 JSON strings,
nil 代表 JSON null.
*/
type Person struct {
	Name string
	Age  int
}

/*
针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:
字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出
到JSON的时候会把该字段对应的值转换成JSON字符串
*/
type Person1 struct {
	Id     int    `json:"-"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school,string"` //会进行二次JSON编码
	Desc   string `json:"desc,omitempty"`
}

func main() {

	s := []byte(`{"name":"Jack","age":34}`)
	var i interface{}
	err := json.Unmarshal(s, &i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	m := i.(map[string]interface{}) //通过断言的方式
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	fmt.Println("---------------")
	p := Person{}
	/*
		首先查找tag含有Foo的可导出的struct字段(首字母大写)
		其次查找字段名是Foo的导出字段
		最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段
	*/
	json.Unmarshal(s, &p)
	fmt.Println(p)

	fmt.Println("----------------")
	b, _ := json.Marshal(p)
	os.Stdout.Write(b)
	fmt.Println("----------------")
	p1 := Person1{Id: 123, Name: "TOm", Age: 13, School: "school1", Desc: ""}
	fmt.Println(p1)

	b, _ = json.Marshal(p1)
	os.Stdout.Write(b)
}

/*
Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
Channel, complex和function是不能被编码成JSON的
嵌套的数据是不能编码的，不然会让JSON编码进入死循环
指针在编码的时候会输出指针指向的内容，而空指针会输出null
*/
