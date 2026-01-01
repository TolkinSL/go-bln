package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 int
	x1 = math.MaxInt
	fmt.Println(x1)
	x1++
	fmt.Println(x1)

	var x2 int8
	x2 = math.MaxInt8
	fmt.Println(x2)
	x2++
	fmt.Println(x2)
	x2--
	fmt.Println(x2)
}
