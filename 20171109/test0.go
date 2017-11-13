package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//解析xmll
//func Unmarshal(data []byte, v interface{}) error
//data接收的是XML数据流
//v是需要输出的结构，定义为interface,支持struct、slice和string
//输出xml
//func Marshal(v interface{}) ([]byte, error)
//func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) 会增加前缀和缩进

//为了正确解析，go语言的xml包要求struct定义中的所有字段必须是可导出的（即首字母大写）
type Person struct {
	XMLName xml.Name `xml:"person"` //struct的一个特性，它们被称为 struct tag,用来辅助反射的
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

type PersonList struct {
	XMLName xml.Name `xml:"persons"`
	Count   int      `xml:"count,attr"` //指定属性
	Person  []Person `xml:"person"`
	desc    string   `xml:",innerxml"`
	//struct的一个字段是string或者[]byte类型且它的tag含有",innerxml"，
	//Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上

	//tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。
	//tag定义了"-",那么不会为该字段解析匹配任何xml数据。
	//tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
	//XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上
}

func main() {

	file, err := os.Open("persons.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	v := PersonList{}
	err = xml.Unmarshal(data, &v) //采用了反射来进行数据的映射，所以v里面的字段必须是导出的
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

	p := &PersonList{Count: 2}
	p.Person = append(p.Person, Person{Name: "aaaa", Age: 233})
	p.Person = append(p.Person, Person{Name: "bbb", Age: 223})
	out, err := xml.MarshalIndent(p, "    ", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(out)
}

/*
设置struct 中字段的tag信息以控制最终xml文件的生成
XMLName不会被输出
tag中含有"-"的字段不会输出
tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。
tag中含有",chardata"，输出为xml的 character data而非element。
tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程
tag中含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
tag中含有"omitempty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string
tag中含有"a>b>c"，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`

	<name>
	<first>Asta</first>
	<last>Xie</last>
	</name>
*/
