// Constants
package main

import (
        "fmt"
)

func main() {
	const pi float32 = 3.14
	fmt.Println(pi)
	var piVar float32 = 3.14
	fmt.Println(piVar)
	piVar = 2
	fmt.Println(piVar)
	fmt.Println(pi)
}