// pointers
package main

import (
	"fmt"
)

// pass by-reference
func addTax(totalPtr *float64)  {
	*totalPtr = *totalPtr * 1.1
	fmt.Println("memory address of totalPtr:", &totalPtr)
	fmt.Printf("total with tax: $%.2f\n", *totalPtr)
}

func main() {

	var total float64 = 5
	fmt.Println("memory address of total:", &total)
	var p *float64 = &total
	fmt.Println("memory address pointed to by p:", p)
	addTax(&total)
	fmt.Printf("total with tax: $%.2f\n", total)

}