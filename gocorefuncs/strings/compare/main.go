package main

import (
	"fmt"
	"strings"
)

func main() {

	// Compare returns an integer comparing two strings lexicographically.
	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	// SP!"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
	fmt.Println(`strings.Compare("a", "b"):`, strings.Compare("a", "b"))
	fmt.Println(`strings.Compare("b", "a"):`, strings.Compare("b", "a"))
	fmt.Println(`strings.Compare("a", "a"):`, strings.Compare("a", "a"))

	/*
		// return true if the same
		fmt.Println(`"a" == "b":`, "a" == "b")
		fmt.Println(`"b" == "a":`, "b" == "a")
		fmt.Println(`"a" == "a":`, "a" == "a")
	*/

	/*
		// greater than operator returns true if left value comes after lexicographically
		// a is 0061 and b is 0062
		fmt.Println(`"a" > "b:"`, "a" > "b")
		fmt.Println(`"b" > "a:"`, "b" > "a")
	*/

	/*
		// lesser than operator returns true if left value comes before lexicographically
		fmt.Println(`"a" < "b":`, "a" < "b")
		fmt.Println(`"b" < "a":`, "b" < "a")
	*/

	/*
		// 5 is 0035 and nine is 0039
		fmt.Println(`"50" < "9":`, "50" < "9")
		// looks at first letter only
		fmt.Println(`"zaaaaaaaa" < "azzzzzzz":`, "zaaaaaaaa" < "azzzzzzz")
	*/

	/*
		// strings.Compare is not the recommended way due to being less efficient
		// strings.Compare does a three way comparison, may be overkill if you just want to make one comparison
		// returning -1, 0, and 1 is not terribly useful, for instance sorting in go returns a bool with Less(i, j int) bool,
		// converting to true or false would be another comparison
		switch strings.Compare("a", "b") {
		case -1:
			fmt.Println("a is less than b")
		case 0:
			fmt.Println("a is equal to b")
		case 1:
			fmt.Println("a is greater than b")
		}
	*/
}
