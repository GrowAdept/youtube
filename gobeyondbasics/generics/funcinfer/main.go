package main

import (
	"fmt"
)

func main() {
	var a int = 2
	var b int8 = 5
	fmt.Println(min(a, b))
}

func min[T int | int8](x, y T) T {
	if x < y {
		return x
	}
	return y
}
