package main

import (
	"fmt"
)

func main() {
	a := min[int8](8, -3)
	b := min(8, -3)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
}

func min[T int | int8](x, y T) T {
	if x < y {
		return x
	}
	return y
}
