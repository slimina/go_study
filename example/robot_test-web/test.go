package main

import (
	"fmt"
	"regexp"
)

func main() {
	someString := "1 test total, 0 passed, 1 failed"
	digitsRegexp := regexp.MustCompile(`.*(\d+) test total, (\d+) passed, (\d+) failed`)
	arr := digitsRegexp.FindStringSubmatch(someString)
	fmt.Println(arr)
}
