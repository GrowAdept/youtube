package main

import "fmt"

func main() {
	//int8: integers (-128 to 127)
	var a int8 = 10 + 10 // no overlow error
	// var a int8 = 100 + 100  // overflow error
	fmt.Println(a) // constant 200 overflows int8
}
