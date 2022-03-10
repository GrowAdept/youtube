package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 100
var wg = sync.WaitGroup{}

func printA() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("A")
	}
	// each of the goroutines runs and calls wg.Done when finished
	wg.Done()
}

func printB() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("B")
	}
	wg.Done()
}

func printC() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("C")
	}
	wg.Done()
}

func printD() {
	for i := 0; i <= count; i++ {
		time.Sleep(time.Nanosecond * 1)
		fmt.Print("D")
	}
	wg.Done()
}

func main() {
	// add 4 to the WaitGroup counter
	wg.Add(4)
	go printA()
	go printB()
	go printC()
	go printD()
	// blocks until the WaitGroup counter is zero
	wg.Wait()
}
