package main

import "fmt"

func main() {
	data1 := []int{}
	data2 := make([]int, 0)

	fmt.Printf("%p\n", data1)
	fmt.Printf("%p\n", data2)

	data1 = append(data1, 1)
	data2 = append(data2, 1)

	fmt.Printf("%p\n", data1)
	fmt.Printf("%p\n", data2)

	arr1 := []int{1, 2, 3, 4, 5}
	var arr2 []int
	arr3 := make([]int, 3)

	copy(arr2, arr1)
	copy(arr3, arr1)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
