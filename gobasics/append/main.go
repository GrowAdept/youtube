// append
package main

import (
	"fmt"
)

func main() {

	// Append: add to end (push)
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)
	// append(s S, x ...T) S
	numbers = append(numbers, 4)
	fmt.Println(numbers)
	// append multiple values
	// numbers = append(numbers, 5, 6, 7)
	// fmt.Println(numbers)
	// append slice
	// s := []int{8, 9, 10}
	// numbers = append(numbers, s...)
	// fmt.Println(numbers)

	// overlapping slice
	/*
		letters := []string{"a", "b", "c", "d"}
		letters2 := append(letters[:3], letters[1:]...)
		fmt.Println(letters2)
	*/

	// remove last of slice (pop)
	/*
		letters := []string{"a", "b", "c", "d"}
		letters = letters[:len(letters)-1]
		fmt.Println(letters)
	*/

	// remove first of slice (shifting)
	/*
		letters := []string{"a", "b", "c", "d"}
		letters = letters[1:]
		fmt.Println(letters)
	*/

	// Prepend: add to start of slice (unshift)
	/*
		letters := []string{"a", "b", "c", "d"}
		letters = append([]string{"z"}, letters[:]...)
		fmt.Println(letters)
	*/

}
