package main

import (
	"fmt"

	"github.com/GrowAdept/youtube/gobeyondbasics/calendar"
)

func main() {
	fmt.Println("hello world")
	calendar.Hello()

	// access through getter and setter methods
	var d calendar.Date
	d.SetYear(2022)
	fmt.Println(d.Year())

	// runs error checking

	// unexported field stop users from assigning directly
	/*
		fmt.Println(d.year)
		// d.year = 2021
		// fmt.Println(d.Year())
	*/

	/*
		// Accidentally exported field
		d.SetDay(22)
		fmt.Println(d.GetDay())
		// error checking is avoided by assigning directly
		d.Day = 100
		fmt.Println(d.Day)
	*/

	// .DisplayWorldDate runs method .printMessage (package scope)
	d.SetMonth(4)
	d.SetDay(10)
	d.DisplayWorldDate()
	// .printMessage method has package scope in calendar.go only
	// d.printMessage("this will not run")

}

type Date struct {
	Year  int
	month int
	day   int
}

func (d *Date) SetYear(y int) {
	d.Year = y
}
