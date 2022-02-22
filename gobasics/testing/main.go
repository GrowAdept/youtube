package main

import (
	"fmt"

	"github.com/GrowAdept/youtube/gobasics/testing/world"
)

func main() {
	fmt.Println(Hello())
	fmt.Println(world.World())
	num := 2
	result := add2(num)
	fmt.Println("result:", result)
}

func Hello() string {
	// return "rock" // code error for demonstration
	return "Hello"
}

func add2(num int) int {
	// num = num + 3 // code error for demonstration
	num = num + 2
	return num
}
