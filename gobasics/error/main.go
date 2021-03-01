// Errors
package main

import (
	"fmt"
	"errors"
)

func rectArea(len float64, wid float64) (area float64, err error) {
	if len < 0 || wid < 0 {
		err = errors.New("cannot not use negative number to calculate area")
		return
	}
	area = len * wid
	return
}

func main() {
	a, err := rectArea(10, 15)
	fmt.Println("Error returned:", err)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("area:", a)
	}
}