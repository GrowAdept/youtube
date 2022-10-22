package main

import (
	"fmt"
	"log"
)

func main() {
	var a string = "dog"
	var b string = "car"
	fmt.Println(min(a, b))
}

func min(x, y interface{}) interface{} {
	if fmt.Sprintf("%T", x) != fmt.Sprintf("%T", y) {
		log.Fatal("not the same type")
	}
	switch x.(type) {
	case int:
		if x.(int) < y.(int) {
			return x
		}
		return y
	case int8:
		if x.(int8) < y.(int8) {
			return x
		}
		return y
	case int16:
		if x.(int16) < y.(int16) {
			return x
		}
		return y
	case int32:
		if x.(int32) < y.(int32) {
			return x
		}
		return y
	case int64:
		if x.(int64) < y.(int64) {
			return x
		}
		return y
	default:
		log.Fatal("Type is unknown!")
		return x
	}
}
