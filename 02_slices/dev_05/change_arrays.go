package main

import "fmt"

type Balance struct {
	blnc int
}

func main() {
	arr := [...]*Balance{
		{blnc: 3},
		{blnc: 6},
		{blnc: 9},
	}

	fmt.Printf("%+v\n", arr[0])

	//v — копия указателя, а сам указатель указывает на оригинальную структуру
	// изменения меняют оригинальные структуры
	for _, v := range arr {
		v.blnc += 1
	}

	fmt.Printf("%+v\n", arr[0])

	fmt.Println("--------")
	arr1 := [...]Balance{
		{blnc: 5},
		{blnc: 3},
		{blnc: 8},
	}

	fmt.Printf("%+v\n", arr1)

	for i, _ := range arr1 {
		arr1[i].blnc += 1
	}

	fmt.Printf("%+v\n", arr1)
}
