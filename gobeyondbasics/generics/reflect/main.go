package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int8 = 3
	findType(a)
}

func findType(a interface{}) {
	fmt.Println(reflect.TypeOf(a))
}
