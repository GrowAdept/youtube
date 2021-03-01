// pointers
package main

import (
	"fmt"
)

// pass by copy
func addTax(t float64) {
	t = t * 1.1
	fmt.Println("memory address of t:", &t)
	fmt.Printf("total with tax: $%.2f\n", t)
}

func main() {
	/*
		var total float64 = 5
		fmt.Println("memory address of total:", &total)
		addTax(total)
		fmt.Printf("total with tax: $%.2f\n", total)
	*/

	// * is the dereference operator
	// & is the address operator
	var a string = "cat"
	var b *string = &a
	fmt.Println("memory address of a:", &a)
	fmt.Println("memory address of b:", &b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("b:", *b)
	// fmt.Println("*a:", *a) // * operator can only be run on pointers

}
