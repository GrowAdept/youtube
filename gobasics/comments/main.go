/*
Package comments is a teaching example for using comments

This package shows
	-single line comments
	-multi-line comments
	-common placements for comments to create go doc for package
*/
package main

import (
	"fmt"
)

// prints passed in value twice
func printTwice(i interface{}) {
	// prints first time
	fmt.Println(i)
	// prints second time
	fmt.Println(i)
}

func main() {
	myVar := "dog"
	printTwice(myVar)
	/*
		var myVar int = "cat"
		printTwice(myVar)
	*/
}
