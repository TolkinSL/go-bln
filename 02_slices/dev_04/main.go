package main

import "fmt"

type account struct {
	balance int
}

func main() {
	{
		acc1 := [...]*account{
			{balance: 300},
			{balance: 400},
			{balance: 250},
		}

		for _, a := range acc1 {
			a.balance += 10
			fmt.Println((*a).balance)
		}
		fmt.Println(acc1)

		acc2 := [...]account{
			{balance: 300},
			{balance: 400},
			{balance: 250},
		}

		for idx := range acc2 {
			acc2[idx].balance += 5
		}
		fmt.Println(acc2)
	}
}
