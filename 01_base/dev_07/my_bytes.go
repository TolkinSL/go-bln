package main

import "fmt"

func main() {
	var x1 int8 = 0b0011
	var y1 int8 = 0b0101
	var z1 int8 = 0b0110

	fmt.Printf("%08b\n", x1)
	fmt.Printf("%08b\n", y1)
	fmt.Println("--------")
	fmt.Printf("%08b\n", x1 & y1)
	fmt.Printf("%08b\n", x1 | y1)
	fmt.Printf("%08b\n", ^x1) //инверсия меняет знак 1-й бит оказывается 1
	fmt.Printf("%08b\n", x1 << 2)
	fmt.Println("--------")
	fmt.Printf("%08b\n", x1 & 1)
	fmt.Printf("%08b\n", y1 & 1)
	fmt.Printf("%08b\n", z1 & 1)

}
