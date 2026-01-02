package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var z int = 100
	fmt.Println("z -", unsafe.Sizeof(z))

	var r rune = 'R'
	fmt.Println("r -", unsafe.Sizeof(r))

	// x := new(int)
	y := new(int)

	// ptrX := unsafe.Pointer(x)
	ptrY := unsafe.Pointer(y)
	addressY := uintptr(unsafe.Pointer(y))

	fmt.Printf("y - %#v\n", y)
	fmt.Printf("ptrY - %#v\n", ptrY)
	fmt.Printf("addY - %#v\n", addressY)

}
