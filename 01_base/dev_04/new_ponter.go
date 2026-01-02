package main

import "fmt"

func myTest(p **int) {
	var y int = 300
	fmt.Println("myTest")
	fmt.Printf("&y - %#v\n", &y)
	fmt.Printf("p - %#v\n", p)
	*p = &y
	fmt.Printf("p - %#v\n", p)
}

func main() {
	var x int = 100
	var p *int = &x

	fmt.Printf("x - %#v\n", x)
	fmt.Printf("p - %#v\n", p)

	myTest(&p)
	fmt.Printf("x - %#v\n", x)
	fmt.Printf("*p - %#v\n", *p)
}
