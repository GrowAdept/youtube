package calendar

import (
	"errors"
	"fmt"
)

// notice Date is capitalized and visible out side of package
// notice fields are not capitalized and are not visible outside of package (package scope)
type Date struct {
	year  int
	month int
	Day   int
}

// Methods are capitalized and visible outside of package (universe scope)
// Setter methods require * to change fields
// Go convention for naming setter methods is SetFieldname
// Setter methods can be used to force validation of data
func (d *Date) SetYear(y int) error {
	if y < 0 {
		return errors.New("negative value for year")
	}
	d.year = y
	return nil
}

func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("value for month less than 1 or greater than 12")
	}
	d.month = m
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("value for day less than 1 or greater than 31")
	}
	d.Day = day
	return nil
}

// If any methods uses a * then it's idomatic in Go to use * in all that types methods
// Go convention for naming getter methods is to just use field name, do not use Get
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) GetDay() int {
	return d.Day
}

func (d *Date) DisplayWorldDate() {
	d.printMessage("DisplayWorldDate method running")
	fmt.Printf("%v-%v-%v", d.year, d.month, d.Day)
}

// unexported method can be run inside of local package but not outside
func (d *Date) printMessage(s string) {
	fmt.Println(s)
}
