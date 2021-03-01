// structs
package main

import (
	"fmt"
)

var carlFirstName = "Carl"
var carlLastName = "Jefferson"
var carlID = 9
var carlSalary = 55000
var carlAdress = "200 N Pine St Austin, TX 77731"
var carlPhoneNumber = 4441118888

type employee struct {
	firstName   string
	lastName    string
	id          int
	salary      float64
	address     string
	phoneNumber int
}

func main() {
	john := employee{
		firstName:   "John",
		lastName:    "Smith",
		id:          10,
		salary:      50000,
		address:     "2200 E Main Austin, TX 77731",
		phoneNumber: 4446663333,
	}
	fmt.Println(john)
	// fmt.Println(carlFirstName, carlLastName, carlID, carlSalary, carlAdress, carlPhoneNumber)
	// fmt.Println(john.lastName)
	// ben := employee {firstName: "Ben", lastName: "Johnson"}
	// fmt.Println(ben)
	// ben.id = 12
	// fmt.Println(ben)

	/*
		var jane employee
		fmt.Println(jane)
		jane.firstName = "Jane"
		jane.lastName = "Swanson"
		jane.id = 13
		fmt.Println(jane)
	*/

	// megan := employee{"Megan", "Wilson", 14, 60000, "500 E 4th Austin, TX 77731", 2227774444}
	// fmt.Println(megan)
}
