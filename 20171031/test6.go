package main

import (
	"fmt"
	"math"
)

//method的语法,在func后面增加了一个receiver(也就是method所依从的主体)
//func (r ReceiverType) funcName(parameters) results
//Receiver可以是值传递，也可以是引用传递， 两者的差别在于, 指针作为Receiver会对实例对象的内容发生操作,
//而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。

/*
在使用method的时候重要注意几点
1.虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
2.method里面可以访问接收者的字段
3.调用method通过.访问，就像struct里面访问字段一样
*/
type Rectangle struct { //矩形
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) arae() float64 {
	r.width = 0 //值传递，内部修改值，不会修改原实例的属性
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//method 以定义在任何你自定义的类型、内置类型、struct等各种类型上面
const ( //枚举
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color int //自定义类型
type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box // 一个slice

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

//引用传递
func (b *Box) SetColor(c Color) {
	b.color = c
}

//查找体积最大的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK) //receiver是指针，他自动帮你转，不是要(&bl[i]).SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {

	r := Rectangle{10.2, 20.23}
	var c = Circle{radius: 10.2}
	fmt.Println(r.arae())
	fmt.Println(c.area())
	fmt.Println(r)

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Println(boxes)
	fmt.Println(len(boxes))

	fmt.Println(boxes.BiggestColor()) //默认调用String方法
	boxes[5].SetColor(RED)
	fmt.Println(boxes.BiggestColor().String())
}
