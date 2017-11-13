package main

import (
	"fmt"
	json "github.com/bitly/go-simplejson"
)

type Person struct {
	Id     int    `json:"-"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school,string"` //会进行二次JSON编码
	Desc   string `json:"desc,omitempty"`
}

func main() {
	js, _ := json.NewJson([]byte(`{"s1":{"array":[1,"2",3],"string": "simplejson","bool": true,"float": 5.150}}`))

	arr, _ := js.Get("s1").Get("array").Array()
	fmt.Println(arr)
	f, _ := js.Get("s1").Get("float").Float64()
	ms := js.Get("s1").Get("string").MustString()
	fmt.Println(f)
	fmt.Println(ms)
	fmt.Println(js.GetPath("s1", "array"))

	js.Get("s1").Del("array")

	fmt.Println(js)
	js.Set("s2", "hellp")
	fmt.Println(js.Map())

}
