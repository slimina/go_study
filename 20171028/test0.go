package main

//. "fmt" 省略引用
//out "fmt" 别名引用
import (
	"fmt"
)

//首字母大写为public，小写为private
var name = "jack"

var (
	a = 1
	b = 2
)

//常量定义
const PI = 3.14
const (
	str1 = "hello"
	str2 = "world"
)

//一般类型声明
type newType int

type (
	x   int  //取决于运行平台 32或64位
	x1  uint //无符号整型
	x8  int8 //一个字段 -128~127 / 0-255 （uint8） 还有 int16,int32.int64
	x32 rune //int32别名
	y   string
	byt byte    // uint8别名
	bl  bool    //一个字节  true or false
	fl  float32 //4字节，小数位精确到7位，float64 8字节，小数位精确到15位
/*
除此之外，还有
复数类型 complex64、complex128
uintptr 足够保存指针的32或64位整型
其他值类型: array struct  string
引用类型：slice  map  chan
接口类型：interface
函数类型：func
默认值:值类型的都是0，bool为false，string为空字符串
*/
)

//类型别名
type (
	bieming string
)

var s1 bieming

//结构体声明
type gopher struct{}

//接口声明
type golang interface{}

//入口函数
func main() {
	fmt.Printf("你好，%s", name)
	fmt.Println()
	fmt.Printf("a+b=%d", a+b)
	fmt.Println()
	s1 = "别名类型"
	fmt.Printf("你好，%s", s1)
	fmt.Println()
}
