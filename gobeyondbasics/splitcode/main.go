package main

import (
	"fmt"
	thing "splitcode/doathing"
	part "splitcode/doathing/doapart"
)

func init() {
	fmt.Println("4 pkg: main     file: main.go     msg: this is an init function")
}

func main() {
	fmt.Println("6 pkg: main     file: main.go     msg: first func in func main()")
	printProverb()
	thing.PrintProverb(thing.Proverb2)
	// thing.PrintProverb(thing.proverb3) // not exported
	part.PrintAnotherProv(part.P)
}
