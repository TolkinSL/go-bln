package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int32 = 0x12345678
	pointer := unsafe.Pointer(&num)

	fmt.Printf("%#v\n", &num)
	fmt.Printf("%#v\n", pointer)

	for i := 0; i < 4; i++ {
		byteValue := unsafe.Add(pointer, i)
		fmt.Printf("bv - %#x\n", *(*int8)(byteValue))
	}

	var num2 int16 = 0x0001
	fmt.Printf("num2 - %#x\n", *(*int8)(unsafe.Pointer(&num2)))
}
