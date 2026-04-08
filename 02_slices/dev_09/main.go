// go build -gcflags="-m"

package main

import (
)

func main() {

	// Срезы до 64Кб остаются в стеке
	
	slice1 := make([]byte, 64<<10)
	_ = slice1

	slice2 := make([]byte, 64<<10+1)
	_ = slice2
}
