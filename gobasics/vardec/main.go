package main

import (
        "fmt"
)

func main() {
	// all 3 methods require you to know your data type
	// use this when you already know your value
	var a int8 = 3
	// use this when you don't know your value yet
	var b int16
	// use this when you are just trying to get something running (not optimized)
	c := 127
	b = 12

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Printf("data type of a: %T\n", a)
	fmt.Printf("data type of b: %T\n", b)
	fmt.Printf("data type of c: %T\n", c)
}