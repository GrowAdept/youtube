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

	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(letters)
	var s []string = letters[2:6]
	fmt.Println(s)
	s[3] = "z"
	fmt.Println(letters)
	fmt.Println(s)
}
