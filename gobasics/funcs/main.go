// Functions
package main

import (
	"fmt"
)

func doSomething() {
	fmt.Println("our function did something")
}

func rectArea(len float64, wid float64) (area float64) {
	area = len * wid
	return
}

func main() {
	doSomething()
	a := rectArea(10, 15)
	fmt.Println("area:", a)
}