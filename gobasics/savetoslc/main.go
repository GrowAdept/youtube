// slices provides access to their underlying array
package main

import (
	"fmt"
)

func main() {

	/*
	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
						|    |    |    |    |    |    |   |
					    0    1    2    3    4    5    6   7    
	*/
	/*
	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
    fmt.Println(letters)
    var s []string = letters[2:6]
	fmt.Println(s)
	s[3] = "z"
	fmt.Println(letters)
	fmt.Println(s)
	*/

	
	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
    fmt.Println(letters)
	// make(type, length, capacity)
	var s = make([]string, 4)
	fmt.Println(s)
	// copy(dst, src []T) int
	num := copy(s, letters[2:6])
	fmt.Println("num:", num)
	fmt.Println(s)
	s[3] = "z"
	fmt.Println(letters)
	fmt.Println(s)
	
}