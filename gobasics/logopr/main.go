// Logical Operators
package main

import (
	"fmt"
)

func main() {

	// logical AND
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	/*
		// car rental example
		hasLicense := false
		hasPayment := true
		if hasLicense && hasPayment {
				fmt.Println("eligible to rent car")
		} else {
				fmt.Println("not eligible to rent car")
		}
	*/

	/*
		// logial OR
		fmt.Println(true || true)
		fmt.Println(true || false)
		fmt.Println(false || true)
		fmt.Println(false || false)
	*/

	/*
		// car purchase example
		dwnPymnt := 2000
		creditScr := 550
		if dwnPymnt > 4000 || creditScr >= 500 {
				fmt.Println("eligible to make payments")
		} else {
				fmt.Println("not eligible")
		}
	*/

	/*
		// logical NOT
		a := true
		b := false
		fmt.Println(a)
		fmt.Println(!a)
		fmt.Println(b)
		fmt.Println(!b)
	*/

}
