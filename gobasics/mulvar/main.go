// Declaring Multiple Variables
package main

import (
        "fmt"
)

func main() {
	tree1 := "oak"
	tree2 := "maple"
	tree3 := "elm"
	tree4 := "pine"

	var tree5, tree6, tree7, tree8 string = "oak", "maple", "elm", "pine"

	fmt.Println("tree1:", tree1, "tree2:", tree2, "tree3:", tree3, "tree4:", tree4)
	fmt.Println("tree5:", tree5, "tree6:", tree6, "tree7:", tree7, "tree8:", tree8)
}