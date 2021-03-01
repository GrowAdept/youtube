// Slices
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
	/*
		s =  [c  d  e  f]
			 |  |  |  |  |
		     0  1  2  3  4
	*/
	//fmt.Println("s[1:3]:", s[1:3])
	//fmt.Println(letters[:4])
	//fmt.Println(letters[2:])
	//fmt.Println(letters[:])

	// make(type, length, capacity)
	/*
		a := make([]int, 4, 8)
		fmt.Println("a:", a)
		b := make([]int, 8, 8)
		fmt.Println("b:", b)
		b = append(b, 1)
		fmt.Println("b:", b)
	*/
}
