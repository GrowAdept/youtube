// Increment and Decrement Statements
package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("a:", a)

	a = a + 1
	fmt.Println("a:", a)
	//  commented out to hide suggestion in editor, uncomment to see editor suggestion
	// a += 1 // linter recommends a++ instead
	fmt.Println("a:", a)
	a++
	fmt.Println("a:", a)
	/*
		a--
		fmt.Println("a:", a)
	*/

	/*
		for i := 0; i < 11; i++ {
			fmt.Println("i:", i)
		}
	*/

}
