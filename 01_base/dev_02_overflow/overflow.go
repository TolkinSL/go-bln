package main

import (
	"errors"
	"fmt"
	"math"
)

func inc1(num int8) (int8, error) {
	if num == math.MaxInt8 {
		return 0, errors.New("Number Overflow")
	}

	myNum := num + 1
	return myNum, nil
}

func add(num1, num2 int8) (int8, error) {
	if num2 > 0 && num1 > math.MaxInt8-num2 {
		return 0, errors.New("Overflow +")
	}

	if num2 < 0 && num1 < math.MaxInt8-num2 {
		return 0, errors.New("Overflow -")
	}

	return num1 + num2, nil
}

func main() {
	fmt.Println("Max int8", math.MaxInt8)
	fmt.Println("Min int8", math.MinInt8)
	var myTest int8 = 126

	res1, err := inc1(myTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res1)
	}

	myTest = 127
	res2, err := inc1(myTest)
	if err != nil {
		fmt.Printf("%v - %#v\n", res2, err)
	} else {
		fmt.Println(res2)
	}

	sum1, errSum := add(126, 1)
	if errSum != nil {
		fmt.Printf("add(126, 1) %v - %#v\n", sum1, errSum)
	} else {
		fmt.Println(sum1)
	}

	sum2, errSum := add(127, 1)
	if errSum != nil {
		fmt.Printf("add(127, 1) %v - %#v\n", sum2, errSum)
	} else {
		fmt.Println(sum1)
	}

	sum3, errSum := add(-128, -1)
	if errSum != nil {
		fmt.Printf("add(-128, -1) %v - %#v\n", sum3, errSum)
	} else {
		fmt.Println(sum1)
	}
}
