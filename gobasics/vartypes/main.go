//Data Types
package main

import (
	"fmt"
)

func main() {
	var a bool = true
	var b string = "truck"
	var c int = 210
	var d float64 = 2.732
	var e complex128 = 5 + 3i

	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of b: %T\n", b)
	fmt.Printf("type of c: %T\n", c)
	fmt.Printf("type of d: %T\n", d)
	fmt.Printf("type of e: %T\n", e)
}
