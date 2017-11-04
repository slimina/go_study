package main

import (
	"fmt"
	"math/rand"
)

//math/rand包提供伪随机数生成。例如，rand.Intn返回一个随机int n，0 <= n <100。
//rand.Float64返回一个float64 f，0.0 <= f <1.0。这可以用于生成其他范围内的随机浮点
func main() {
	//每次输出一样，使用了的相同的种子
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(100), ",")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Float64())
	}
	//要让伪随机数生成器有确定性，可以给它一个明确的种子。
	s1 := rand.NewSource(30)
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		fmt.Print(r1.Intn(100), ",")
	}
	//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
	s2 := rand.NewSource(30)
	r2 := rand.New(s2)
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(r2.Intn(100), ",")
	}
}
