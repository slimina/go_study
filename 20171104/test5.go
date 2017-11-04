package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Page   int
	Fruits []string
}

type Person1 struct {
	Page   int      `json:"page"` //json序列化字段
	Fruits []string `json:"fruits"`
}

//Go提供对JSON编码和解码的内置支持，包括内置和自定义数据类型。
func main() {

	bl, err := json.Marshal(true)
	fmt.Println(string(bl), err)

	intb, _ := json.Marshal(23)
	fmt.Println(string(intb))
	strB, _ := json.Marshal("gopher")

	fmt.Println(string(strB))
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	fmt.Println("=================")
	res1 := &Person{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1)
	fmt.Println(string(res1B))

	res2 := &Person1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2)
	fmt.Println(string(res2B))

	fmt.Println("=====================")

	//反序列化
	bty := []byte(`{"name":"Tome","age":24,"strs":["abc","def"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(bty, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	fmt.Println(dat["age"].(float64))

	fmt.Println(dat["strs"].([]interface{})) //访问嵌套数据需要先强制转换为inferface
	fmt.Println(dat["strs"].([]interface{})[0].(string))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Person1{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	fmt.Println("================")
	enc := json.NewEncoder(os.Stdout) //输出控制台
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
