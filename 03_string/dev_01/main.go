package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func foo(s string) {
	s = "Привет"
}

func main() {
	str1 := "Hello\n World!" //Интерпретированая
	str2 := `Hello\n World!` //Не Интерпретированая

	fmt.Println(str1)
	fmt.Println(str2)

	var builder strings.Builder
	str3 := "Hello"
	builder.WriteString(str3)
	str4 := builder.String()
	fmt.Println(str4)

	str5 := "Hello"
	str6 := "Пello"
	fmt.Println(str5, str6)
	fmt.Println(len(str5), len(str6))
	fmt.Println("utf8.Rune:", utf8.RuneCountInString(str6))

	newStr1 := str5[1:4]
	fmt.Println(newStr1)

	str7 := "Привет"
	newStr2 := str7[1:4]
	fmt.Println(newStr2)

	sliceStr := []byte(str5)
	fmt.Println(str5, string(sliceStr))

	str8 := "Hello Привет"
	for _, symbol := range str8 {
		fmt.Println(symbol, string(symbol))
	}

	foo(str8)
	fmt.Println(str8)

	myHello := []byte("Hello")
	myHello = append(myHello, str7...)
	fmt.Println(string(myHello))

}