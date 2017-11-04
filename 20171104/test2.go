package main

import (
	"fmt"
	"sort"
)

//sort 包实现了内置和用户定义类型的排序。排序方法特定于内置类型;
//排序是就地排序，因此它会更改给定的切片，并且不返回新的切片。
//可以使用sort来检查切片是否已经按照排序顺序。
func main() {
	ss := []string{"f", "b", "c", "i", "a", "e"}
	sort.Strings(ss)
	fmt.Println(ss)

	n := []int{4, 7, 1, 5, 2, 8, 3}
	sort.Ints(n)
	fmt.Println(n)

	fmt.Println(sort.IntsAreSorted(n))

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

//自定义排序
//需要实现sort.Interface - Len，Less和Swap - 在这个类型上，
//所以可以使用 sort 包中的一般Sort函数。Len和Swap通常在类型之间是相似的，Less保存实际的自定义排序逻辑。

type ByLength []string //自定义类型，按照字符串长度排序
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
