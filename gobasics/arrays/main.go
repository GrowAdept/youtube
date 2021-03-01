// Arrays
package main

import (
	"fmt"
)

func main() {

	// real world examples: pages in a book, post office boxes, lockers

	// array initialization example
	numbers := [4]int{33, 18, 2, 9}
	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	floats := [3]float64{3.11, 9, 10.762}
	booleans := [...]bool{true, true, false, true, false}
	fmt.Println(numbers)
	fmt.Println(letters)
	fmt.Println(floats)
	fmt.Println(booleans)

	/*
		// array declaration example
		var numbers [4]int
		var letters [7]string
		var floats [3]float64
		var booleans [4]bool
		fmt.Println(numbers)
		fmt.Println(letters)
		fmt.Println(floats)
		fmt.Println(booleans)
	*/

	/*
			["a", "b", "c", "d", "e", "f", "g"]
		index 0    1    2    3    4    5    6
	*/
	/*
		letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
		// access by index
		fifth := letters[4]
		fmt.Println(fifth)
		// access first element in array
		first := letters[0]
		fmt.Println(first)
		// access last element in array
		length := len(letters)
		fmt.Println("length:", length)
		last := letters[length-1]
		fmt.Println(last)
		// last in one line
		fmt.Println(letters[len(letters)-1])
	*/

	/*
		// locker example
		lockers := [6]string{"Bill", "Jane", "Becky", "Mark", "Logan", "Jessica"}
		// Jane and Logan want to switch lockers
		student1 := lockers[1]
		student2 := lockers[4]
		lockers[1] = student2
		lockers[4] = student1
		fmt.Println(lockers)
	*/

}
