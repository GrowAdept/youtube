package main

import (
	"fmt"
)

func main() {
	// once a variable's data type is declared, we cannot change it's data type to something else
	var age int = 25
	fmt.Println("age:", age)
	// one year later
	age = 26
	fmt.Println("age:", age)
}
