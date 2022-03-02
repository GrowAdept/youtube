package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/GrowAdept/youtube/gobeyondbasics/calendar"
)

type date struct {
	day int
}

func (d *date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("value for day less than 1 or greater than 31")
	}
	d.day = day
	return nil
}

func (d *date) Day() int {
	return d.day
}

func main() {
	var myDate date
	myDate.day = 50
	fmt.Println("myDate day:", myDate.day)

	// access through getter and setter methods
	var d calendar.Date
	err := d.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("year:", d.Year())

	// unexported field stop users from assigning directly

	// fmt.Println("year:", d.year)
	// d.year = 2021
	fmt.Println("year:", d.Year())

	/*
		err = d.SetDay(50)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// Accidentally exported field
	d.SetDay(22)
	fmt.Println("day:", d.GetDay())
	// error checking is avoided by assigning directly
	d.Day = 100
	fmt.Println("day:", d.Day)

	// .DisplayWorldDate runs method .printMessage (package scope)
	d.SetMonth(4)
	d.SetDay(10)
	d.DisplayWorldDate()
	// .printMessage method has package scope in calendar.go only
	// d.printMessage("this will not run")

}
