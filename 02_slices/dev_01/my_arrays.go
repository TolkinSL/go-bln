package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr1 := [10]int{}
	fmt.Printf("%#v\n", arr1)
	fmt.Println(cap(arr1))

	arr2 := [3]string{"a", "b", "c"}
	arr3 := [3]string{"a", "b", "c"}
	arr4 := [3]string{"a", "b", "e"}
	fmt.Println(arr2 == arr3) //true
	fmt.Println(arr2 == arr4) //false

	var arr5 [2]int
	fmt.Println("arr5 len",len(arr5))
	fmt.Println("arr5 cap",cap(arr5))
	fmt.Printf("%#v\n", arr5)
	fmt.Println("arr5 unsafe Sizeof",unsafe.Sizeof(arr5))
}
