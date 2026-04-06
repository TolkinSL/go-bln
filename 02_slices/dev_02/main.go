package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [3]int{1, 2, 3}

	const elemSize = unsafe.Sizeof(int(0))
	pointer := unsafe.Pointer(&arr)
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])

	first := *(*int)(unsafe.Add(pointer, elemSize*0))
	fmt.Println(first)

	second := *(*int)(unsafe.Add(pointer, elemSize*1))
	fmt.Println(second)
}
