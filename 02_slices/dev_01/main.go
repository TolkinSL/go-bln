package main

import (
	"fmt"
	"unsafe"
)

func main() {
	{
		var arr1 [5]int
		fmt.Printf("arr1: %d\n", len(arr1))
		var arr2 [2][5]int

		_ = arr1
		fmt.Printf("arr2: %#v\n", arr2)

		arr3 := [...]int{4, 2, 3}
		_ = arr3

		arr4 := [5]int{2: 12, 4: 17}
		fmt.Printf("arr4: %#v\n", arr4)
		fmt.Printf("arr4 %p\n", &arr4[0])

		arr5 := arr4[2:5]
		fmt.Printf("arr5: %#v\n", arr5)

		arr6 := [3]int{1, 2, 3}
		arr7 := [3]int{1, 2, 3}
		fmt.Println("arr6 = arr7 :", arr6 == arr7)
		fmt.Println("arr6 = arr3 :", arr6 == arr3)
	}

	fmt.Println("----------")

	{
		var arr1 [4]byte
		fmt.Println("arr1 size:", unsafe.Sizeof(arr1))

		var arr2 [4]int
		fmt.Println("arr1 size:", unsafe.Sizeof(arr2))
	}
}
