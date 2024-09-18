package main

import (
	"fmt"
)

func main() {
	var str string
	var stroutput string = "0"
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		if str[i] > stroutput[0] {
			stroutput = string(str[i])
		}
	}
	fmt.Println(stroutput)
}
