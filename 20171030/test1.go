package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is error!")
	if err != nil {
		fmt.Println(err)
	}
}
