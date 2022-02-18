/*
// make sure your are in the root of the module for commands
// get package if we don't have it already
go get github.com/common-nighthawk/go-figure
// creates a go.mod file to track dependencies
go mod int <module-path>
// will add / remove dependencies in our code to go.mod and creates go.sum
go tidy
*/
package main

import (
	"fmt"
	// "github.com/common-nighthawk/go-figure"
)

func main() {
	print("Hello World\n")
	fmt.Println("Hello World")

	/*
		myFigure := figure.NewFigure("Hello World", "", true)
		myFigure.Print()
	*/
}
