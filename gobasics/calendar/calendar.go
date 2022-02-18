package calendar

import "fmt"

// notice Date is capitalized and visible out side of package
// notice fields are not capitalized and are not visible outside of package (package scope)
type Date struct {
	year  int
	month int
	day   int
}

func Hello() {
	fmt.Println("Hello")
}

// Methods are capitalized and visible outside of package (universe scope)

// Setter methods require * to change fields
// Go convention for naming setter methods is SetFieldname
// Setter methods can be used to force validation of data
func (d *Date) SetYear(y int) {
	d.year = y
}

func (d *Date) SetMonth(m int) {
	d.month = m
}

func (d *Date) SetDay(day int) {
	d.day = day
}

// If any methods uses a * then it's idomatic in Go to use * in all that types methods
// Go convention for naming getter methods is to just use field name, do not use Get
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}

func (d *Date) DisplayWorldDate() {
	fmt.Printf("%v-%v-%v", d.year, d.month, d.day)
}
