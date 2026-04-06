// go build -gcflags='-m'
package main

func foo() [1024]int8 {
	var arr [1024]int8
	return arr
}

func fooPointer() *[1024]int8 {
	var arrP [1024]int8
	return &arrP
}

func main() {
	var arrStack [1024] int8
	_ = arrStack

	var arrHeap [10 << 20] int8
	_ = arrHeap

	fooArr := foo()
	_ = fooArr
}

/*
# command-line-arguments
./main.go:4:6: can inline foo
./main.go:9:6: can inline main
./main.go:16:15: inlining call to foo
./main.go:13:6: moved to heap: arrHeap
*/