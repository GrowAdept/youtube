// Switch Statements
package main

import "fmt"

func main() {

	n := 4

	if n == 0 {
		fmt.Println("zero")
	} else if n == 1 {
		fmt.Println("one")
	} else if n == 2 {
		fmt.Println("two")
	} else if n == 3 {
		fmt.Println("three")
	} else if n == 4 {
		fmt.Println("four")
	} else if n == 5 {
		fmt.Println("five")
	}

	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	}

	/*
		n := 54

		switch n {
		case 1:
				fmt.Println("one")
		case 2:
				fmt.Println("two")
		case 3:
				fmt.Println("three")
		case 4:
				fmt.Println("four")
		case 5:
				fmt.Println("five")
		default:
				fmt.Println("number not found")
		}
	*/

	/*
			// fallthrough example
			switch "b" {
			case "a":
		    	fmt.Println("a")
		    	fallthrough
			case "b":
		    	fmt.Println("b")
		    	fallthrough
			case "c":
		    	fmt.Println("c")
				fallthrough
			case "d":
				fmt.Println("d")
			}
	*/

	/*
		// no condition example
		var score int = 81
		var grade string

		switch {
			case score >= 90: grade = "A"
			case score >= 80: grade = "B"
			case score >= 70: grade = "C"
			case score >= 60: grade = "D"
			default: grade = "F"
		 }

		 fmt.Println("Grade:", grade)
	*/
}
