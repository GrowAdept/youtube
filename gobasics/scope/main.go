// Scope
package main

import (
	"fmt"
	"gobasics/pets"
)

// global variable
var pet string = "gold fish"

func printPet() {
	fmt.Println("pet from inside printPet():", pet)
	pet := "bird"
	fmt.Println("pet after setting inside printPet():", pet)
}

func main() {
	fmt.Println("pet from 1st line in main():", pet)
	// local variable
	var pet string = "dog"
	fmt.Println("pet after setting var pet in main():", pet)
	printPet()

	{
		fmt.Println("pet inside local block:", pet)
		// variable local to block
		var pet string = "bird"
		fmt.Println("pet after setting var pet in local block:", pet)
	}
	// pets.pet2 is not is not exported since it's not capitalized in package pets
	// fmt.Println("pet2 from pets package:", pets.pet2)
	fmt.Println("pet3 from pets package", pets.Pet3)
}
