package main

import "fmt"

func main() {
	{
		data1 := make([]int, 0, 6)
		data1 = append(data1, []int{1, 2, 3, 4}...)
		data2 := data1[2:]
		fmt.Println("data2 len:", len(data2), "cap:", cap(data2))
		data2[0] = 10
		fmt.Println("data1:", data1)
		fmt.Println("data2:", data2)
		fmt.Println("data1 len:", len(data1), "cap:", cap(data1))
		fmt.Println("data2 len:", len(data2), "cap:", cap(data2),"\n")

		data2 = append(data2, 15)
		fmt.Println("data1:", data1)
		fmt.Println("data2:", data2)
		fmt.Println("data1 len:", len(data1), "cap:", cap(data1))
		fmt.Println("data2 len:", len(data2), "cap:", cap(data2))

		data1 = append(data1, 150)
		fmt.Println("data1:", data1)
		fmt.Println("data2:", data2)
		fmt.Println("data1 len:", len(data1), "cap:", cap(data1))
		fmt.Println("data2 len:", len(data2), "cap:", cap(data2))
	}
}
