package main

import (
	"fmt"
)

//函数作为值、参数
//在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
//函数当做值和类型在我们写一些通用接口的时候非常有用
//定义：type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
type testInt func(int) bool

func isOdd(a int) bool {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice=", slice)

	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)

	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}
