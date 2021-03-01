// Maps
package main

import "fmt"

func main() {

	// map [ KeyType ] ElementType
	phoneBook := map[string]int{}
	fmt.Println(phoneBook)
	phoneBook["John Smith"] = 1112223333
	phoneBook["Jane Doe"] = 4445556666
	fmt.Println(phoneBook["John Smith"])
	fmt.Println(phoneBook)

	/*
			//map literal
		    phoneBook := map[string]int {"John Smith": 1112223333, "Jane Doe": 4445556666}
		    fmt.Println(phoneBook["Jane Doe"])
		    phoneBook["Jane Doe"] = 5555555555
		    fmt.Println(phoneBook["Jane Doe"])
	*/

	/*
		// a map is an unordered group of elements
		for i := 0; i < 21; i++ {
			fmt.Println(phoneBook)
		}
	*/

	/*
		// unitialized map is nil
		var phoneBook map[string]int
		phoneBook["John Smith"] = 1112223333
		fmt.Println (phoneBook)
	*/

	/*
		// make(Type, n)      space for approximately n elements
		populations := make(map[string]int)
		populations["China"] = 1394015977
		populations["India"] = 1326093247
		populations["United States"] = 332639102
		populations["Indonesia"] = 267026366
		populations["Pakistan"] = 233500636
		populations["not a real place"] = 200100333
		fmt.Println(populations)
	*/

	/*
		// delete(m, k)  // remove element m[k] from map m
		delete(populations, "not a real place")
		fmt.Println(populations)
	*/
}
