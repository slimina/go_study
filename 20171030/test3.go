package main

import (
	"fmt"
)

// --- 动态数组
//slice并不是真正意义上的动态数组，而是一个引用类型。
//slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

var (
	slice = []int{1, 2, 3, 3, 4}
)

func main() {
	fmt.Println(slice)
	//slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。
	var arr = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	var a []byte = arr[1:2]
	var b = arr[1:4]
	var c []byte
	c = arr[2:5]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	d := arr[:3]
	fmt.Println(d)
	//slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	e := arr[0:]
	fmt.Println(e)
	//如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
	fmt.Println(arr[:])

	//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，
	/*
		从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
		一个指针，指向数组中slice指定的开始位置
		长度，即slice的长度
		最大长度，也就是slice开始位置到数组的最后位置的长度
	*/
	f := arr[:5]
	f[2] = 'm'
	fmt.Println(f)
	fmt.Println(b, arr)
	//从Go1.2开始slice支持了三个参数的slice
	g := arr[1:3:7] //这个的容量就是7-1,最后一位不能超过数组arr的下标
	fmt.Println(g)
	/*
		对于slice有几个有用的内置函数：
		len 获取slice的长度
		cap 获取slice的最大容量
		append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
		copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数,用于将内容从一个数组切片复制到另一个数组切片。
			如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。
		注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。 但当slice中没有剩余空间（即(cap-len) == 0）时，
		此时将动态分配新的数组空间。返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响。
	*/
	fmt.Println(len(g))
	fmt.Println(cap(g))
	h := append(g, 'u', 'v', 'w', 'x', 'y', 'z', '$', '*')
	fmt.Println(h)
	fmt.Println(cap(g), cap(h))
	var j = []byte{'0', '0', '0', '0', '0', '0'}
	i := copy(j, b) //j内容复制b的3个元素
	fmt.Println(b)
	fmt.Println(i)
	fmt.Println(j)

	//遍历
	for k, v := range j {
		fmt.Println(k, "=>", v)
	}
}
