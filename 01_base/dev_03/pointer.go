package main

import "fmt"

func double(num *int) {
	*num *= 2
}

func main() {
	var x int = 100
	var p *int = &x

	fmt.Printf("%#v\n", p)

	*p = 200
	fmt.Printf("%#v\n", x)

	double(&x)
	fmt.Printf("%#v\n", x)
}
