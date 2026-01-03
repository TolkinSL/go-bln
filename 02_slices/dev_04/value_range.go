package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [...]int{3, 5, 8}

	//В Go > 1.22 value создается каждый раз новая, а раньше одна была переменная
	for idx, value := range arr {
		// value += 1

		//В Go 1.22+ range value создаётся новый объект, если вы берёте адрес и/или модифицируете.
		//uintptr прерывает escape analysis → может оставлять один объект.
		//fmt.Println("idx:", idx, "value addr:", uintptr(unsafe.Pointer(&value)), "value:", value)
		fmt.Println("idx:", idx, "value addr:", unsafe.Pointer(&value), "value:", value)
	}
}
