package main

import "fmt"

type proverb string

func init() {
	fmt.Println("5 pkg: main     file: xyz.go      msg: this is an init function")
}

func printProverb() {
	fmt.Println("7 pkg: main     file: xyz.go      msg:", myProverb)
}
