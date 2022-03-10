package main

import (
	"fmt"
	"time"
)

var count = 100

func printA() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("A")
	}
}

func printB() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("B")
	}
}

func printC() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("C")
	}
}

func printD() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("D")
	}
}

func main() {
	/*
		add "go" keyword to create a new goroutine
		only one goroutine at a time is executing, goroutines don't continue executing in any particular order (fastest)
		goroutines have no return values
	*/
	go printA()
	go printB()
	go printC()
	go printD()
	// when main func ends, so does the goroutines it spun up
	time.Sleep(3 * time.Second)
}
