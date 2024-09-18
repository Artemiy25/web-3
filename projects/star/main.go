package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a)
	for i := 0; i < len(a)-1; i++ {
		b = b + string(a[i]) + "*"
	}
	b = b + string(a[len(a)-1])
	fmt.Println(b)
}
