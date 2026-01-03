package main

//go run -gcflags='-m' ./my_heap.go | grep escape

func main() {
	var arrStack [10<<10]int8
	_ = arrStack

	var arrHeap [10<<20 + 1]int8
	_ = arrHeap
}
