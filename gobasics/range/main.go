// Range
package main

import (
	"fmt"
)

func main() {

	// how not to add example (without loop)
	prices := [10]float64{3.12, .99, 2.87, 9.99, 14.78, .56, 11.01, 4.96, 1.86, 4.10}
	var total float64

	total = total + prices[0]
	total = total + prices[1]
	total = total + prices[2]
	total = total + prices[3]
	total = total + prices[4]
	total = total + prices[5]
	total = total + prices[6]
	total = total + prices[7]
	total = total + prices[8]
	total = total + prices[9]

	fmt.Println("the final total is", total)

	/*
		prices := [10]float64{3.12, .99, 2.87, 9.99, 14.78, .56, 11.01, 4.96, 1.86, 4.10}
		var total float64

		for _, v := range prices {
			total = total + v
		}

		fmt.Println("the final total is", total)
	*/

	/*
		prices := [3]float64{3.12, .99, 2.87}

		var total float64

		for i, v := range prices {
				fmt.Println("index:", i, "value:", v, "total:", total)
				total = total + v
		}

		fmt.Println("final price total:", total)
	*/

}
