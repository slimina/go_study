package hello

import (
	"fmt"
)

const (
	PI = 3.14
)

var Str = "hello"

func init() {
	fmt.Println(Str)
	fmt.Println("hello init ..")
}
