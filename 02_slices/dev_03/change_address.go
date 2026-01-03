package main

import (
	"fmt"
	"unsafe"
)

//Не выделять память при компиляции
//go:noinline
func allocation(index int) byte {
	var data [1<<16]byte
	return data[index]
}

func main() {
	var arr [10]int
	//address := (uintptr)(unsafe.Pointer(&arr)) //Избыточная запись приведения типа C++
	address := uintptr(unsafe.Pointer(&arr))
	fmt.Println("arr address:", address)

	//Увеличение размера стека
	allocation(100)

	address = uintptr(unsafe.Pointer(&arr))
	fmt.Println("arr address:", address)
}