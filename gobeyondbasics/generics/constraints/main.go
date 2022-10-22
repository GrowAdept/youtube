package main

import "golang.org/x/exp/constraints"

func main() {
}

func Order[T []int | []int8 | []int16 | []int32 | []int64 | []uint | []uint8 | []uint16 | []uint32 | []uint64 | []float32 | []float64 | []string](x []T) []T {
	// some code ...
	return x
}

func Order2[T constraints.Ordered](x []T) []T {
	// some code ...
	return x
}
