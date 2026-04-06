package main

// go test -bench=. -benchmem capbench_test.go

import "testing"

func BenchmarkWithoutReservation(b *testing.B) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i :=0; i < b.N; i++ {
		target := make([]int, 0)
		for j := 0; j < len(source); j++ {
			target = append(target, source[j])
		}
	}
}

func BenchmarkWithReservation(b *testing.B) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		target := make([]int, 0, len(source))
		for j := 0; j < len(source); j++ {
			target = append(target, source[j])
		}
	}
}

/*
BenchmarkWithoutReservation-12           7339808               161.9 ns/op           248 B/op          5 allocs/op
BenchmarkWithReservation-12             39894856                29.11 ns/op           80 B/op          1 allocs/op
PASS
*/