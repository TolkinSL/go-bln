// go build -gcflags="-m"

package main

import (
	"fmt"
	// "unsafe"
)

//go:noinline
// func allocation(idx int) byte {
// 	var data [1 << 20]byte

// 	return data[idx]
// }

//go:noinline
func grow(n int) {
	var buf [1024]byte
	if n > 0 {
		grow(n - 1)
	}
	_ = buf
}

func main() {
	arr := [3]int{1, 2, 3}
	_ = &arr
	// fmt.Printf("%p\n", &arr)
	// address1 := unsafe.Pointer(&arr)
	// fmt.Printf("unsafe %p\n",address1)
	
	// allocation(100)
	grow(10000)
	// fmt.Printf("%p\n", &arr)
	// address2 := unsafe.Pointer(&arr)
	// fmt.Printf("unsafe %p\n",address2)

	for _, value := range arr {
		// 1.21 &value тоже самое значение адреса
		// 1.22 создается копия value, &value другой адрес
		fmt.Printf("%p\n", &value)
	}
}
