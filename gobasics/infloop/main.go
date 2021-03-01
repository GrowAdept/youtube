// Infinite Loops
package main

import (
	"fmt"
)

func main() {

	i := 0
	for false {
		i = i + 1
		fmt.Println(i)
	}

	/*
		i := 0
		for true{
			i = i + 1
			fmt.Println(i)
		}
	*/

	/*
		for i := 0; i < 10; i-- {
			fmt.Println(i)
		}
	*/

	/*
		var i int = 0
		for {
			i--
			fmt.Println(i)
			if i <= -10 {
				break
			}
		}
	*/

}
