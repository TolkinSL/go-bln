package main

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
	fmt.Printf("%d Mb\n", m.Alloc/(1024*1024))
}
func main() {
	data := []int{1, 2, 3, 4, 5}
	slice := data[:len(data)-1]
	fmt.Println("data:", data)
	fmt.Println("slice:", slice)
	fmt.Println("data len:", len(data), "cap:", cap(data))
	fmt.Println("slice len:", len(slice), "cap:", cap(slice), "\n")

	fmt.Println("удаление первого элемента")
	data = []int{1, 2, 3, 4, 5}
	slice = data[1:]
	fmt.Println("slice:", slice)
	fmt.Println("slice len:", len(slice), "cap:", cap(slice))
	fmt.Println("data:", data, "\n")

	slice = append(data[:0], data[1:]...)
	fmt.Println("slice:", slice)
	fmt.Println("slice len:", len(slice), "cap:", cap(slice))
	fmt.Println("data:", data, "\n")

	/*
		// чтобы не мутировать data
		// Способ 2: append с новым пустым срезом (не связанным с data)
		newSlice := append([]int{}, data[1:]...)

		// Способ 3: используйте полную копию
		newSlice := append(make([]int, 0, len(data[1:])), data[1:]...)
	*/

	fmt.Println("удаление внутри")
	fmt.Println("data:", data)
	slice = append(data[:2], data[3:]...)
	fmt.Println("slice:", slice)
	fmt.Println("slice len:", len(slice), "cap:", cap(slice))

	// Ресайз len
	data = make([]int, 4, 6)
	fmt.Println("data:", data)
	data = data[:6]
	fmt.Println("data:", data, "\n")

	printAlloc()

	//Копирование с очисткой выделенной памяти
	data = []int{1, 2, 3, 4, 5}
	data = append([]int{}, data[:2]...)
	fmt.Println("data:", data, "data len:", len(data), "cap:", cap(data))
}
