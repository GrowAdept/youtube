// variadic functions
package main

import (
	"fmt"
)

var burger, fries, soda, iceCream = 3.05, 1.44, 1.00, 2.25

func main() {
	totalTwoItems("John", 3.05, 1.44)
	// totalTwoItems("John", 3.05, 1.44, 1.00) // too many arguements

	// func Println(a ...interface{}) (n int, err error)
	fmt.Println("a", "b", "c", "d", "e")

	findTotal("John", 3.05, 2.25, 1.44)
	findTotal("John", burger, burger, soda, iceCream, fries, soda, burger, iceCream)

	// variadic arguements may be be omitted by will result in an empty slice
	findTotal("John") // with no actual value passed in for items, items is nil
	// non-variadic arguements are required, expecting string for first arguement
	// findTotal(burger, fries, soda)

	// findTotal2() has a syntax error: cannot use ... with non-final parameter items
	// findTotal2(burger, burger, soda, iceCream, fries, soda, burger, iceCream, "John") // syntax error in function

	var order = []float64{burger, burger, soda, iceCream, fries}
	// findTotal("John", order) // cannot use order (type []float64) as type float64 in argument to findTotal
	findTotal("John", order...)

}

// by placing an ellipsis (three dots) in front of final parameter item
// items var will hold a slice of passed in float64 arguments
func findTotal(name string, items ...float64) {
	var t float64
	for _, v := range items {
		t += v
	}
	// 10% tax
	t = t * 1.1
	fmt.Printf("%v your total is $%.2f\n", name, t)
}

/*
// WARNING: this will not work since the ... is not on the last parameter in function signature
// throws error: syntax error: cannot use ... with non-final parameter items
func findTotal2(items ...float64, name string) {
	var t float64
	for _, v := range items {
		t += v
	}
	// 10% tax
	t = t * 1.1
	fmt.Printf("%v your total is $%.2f\n", name, t)
}
*/

func totalTwoItems(name string, item1 float64, item2 float64) {
	t := item1 + item2
	t = t * 1.1
	fmt.Printf("%v your total is $%.2f\n", name, t)
}
