package main

import "fmt"

func myfunc(arr []int) {
	arr[0] = 123
}

func main() {
	arr1 := make([]int, 3, 6)
	fmt.Println("arr1:", arr1)

	b := []int{1,2,3}
	arr1 = append(arr1, b...)
	fmt.Println("arr1:", arr1)

	arr2 := make([]int, 0)
	fmt.Println("arr2:", arr2)
	arr2 = append(arr2, 1)
	fmt.Println("arr2:", arr2)

	myfunc(arr1)
	fmt.Println("arr1:", arr1)

	arr3 := [...]int{1, 2, 3}
	// Массив копируется при передаче в range
	for _, val := range arr3 {
		fmt.Println("arr3 copy:", val)
	}

	// Передается ссылка на массив (когда большие массивы)
	for _, val := range &arr3 {
		fmt.Println("arr3 point:", val)
	}

	// Передается ссылка на слайс (когда большие массивы)
	for _, val := range arr3[:] {
		fmt.Println("arr3 slice:", val)
	}
}
