package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const elemSize = unsafe.Sizeof(int16(0))
	fmt.Println("Size int16 bytes", elemSize)

	arr := [...]int16{2, 5, 7, 4}
	pointer := unsafe.Pointer(&arr)

	x0 := *(*int16)(unsafe.Add(pointer, 0 * elemSize))
	fmt.Println("arr[0]", x0)

	x2 := *(*int16)(unsafe.Add(pointer, 2 * elemSize))
	fmt.Println("arr[2]", x2)
}
