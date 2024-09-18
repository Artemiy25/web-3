package main

import (
    "fmt"
	"math"
)

var k float64 = 1296.0
var p float64 = 6.0
var v float64 = 6.0

func M() float64 {
    return p*v
}

func W() float64 {
    m := M()
    return math.Sqrt(k/m)
}

func T() float64 {
    w := W()
    return 6/w
}

func main() {
    t := T()
    fmt.Print(t)
}