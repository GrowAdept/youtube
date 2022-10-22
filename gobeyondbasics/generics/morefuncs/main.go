package main

import "fmt"

func main() {
	var a, b int = 9, 1
	fmt.Println(min(a, b))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func min8(x, y int8) int8 {
	if x < y {
		return x
	}
	return y
}

func min16(x, y int16) int16 {
	if x < y {
		return x
	}
	return y
}

func min32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
