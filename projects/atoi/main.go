package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	var str string = ""
	fmt.Scan(&a)
	for a > 0 {
		b = (a % 10) * (a % 10)
		str = strconv.Itoa(b) + str
		a = a / 10
	}
	fmt.Println(str)
}
