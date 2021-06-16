package main

import (
	"fmt"
	"os"
)

func main() {
	// these will all return an empty string if environmental variables are not set
	// func Getenv(key string) string
	goProverb := os.Getenv("GoProverb")
	fmt.Println("goProverb:", goProverb)
	// GOROOT: path to the root of our Golang API (files created from Go's developers)
	goRoot := os.Getenv("GOROOT")
	fmt.Println("goRoot:", goRoot)
	// GOPATH is the path to the root of our workspace (files we will create)
	goPath := os.Getenv("GOPATH")
	fmt.Println("goPath:", goPath)
}
