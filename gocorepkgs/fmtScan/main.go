// fmt package scan functions
package main

import (
	"fmt"
)

func main() {

	// Scan, Scanf and Scanln read from os.Stdin
	// func Scan(a ...interface{}) (n int, err error)
	var name string
	fmt.Print("What is your name? ")
	n, err := fmt.Scan(&name) // text read from stand input
	fmt.Println("number of items successfully scanned:", n, "error:", err)
	fmt.Println("Nice to meet you", name)

	/*
		var s1 string
		var s2 string
		var s3 string
		// Scan scans text read from standard input,
		// storing successive space-separated values into successive arguments.
		fmt.Scan(&s1, &s2, &s3)
		fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
		// Newlines count as space.
		// try dog cat bird seperated by spaces and then seperated by returns
	*/

	/*
		var item string
		var cost float64
		// Newlines in the input must match newlines in the format
		// only stores space-serparated values into successive arguments, not newlines
		// Scanf restricts input to a defined format
		fmt.Scanf("item:%s cost:%f", &item, &cost)
		// Printf prints in a defined format
		fmt.Printf("%s $%.2f", item, cost)
		// the input "item:car cost:6000" works but "car 6000" does not
	*/

	/*
		var food1 string
		var food2 string
		var food3 string
		// Newlines in the input must match newlines in the format
		// only stores space-serparated values into successive arguments, not newlines
		fmt.Print("What is your 3 favorite foods? ")
		// the input "pizza fries soda" work but "pizza \n fries \n soda" does not
		fmt.Scanln(&food1, &food2, &food3)
		fmt.Print(food1, ", ", food2, ", and ", food3, ", all sound tasty")
	*/

	/*
		// scanning from strings
		// var employee string = "John 33 70000"
		var employee string = "John 33 70000"
		var name string
		var age, salary int
		n, err := fmt.Sscan(employee, &name, &age, &salary)
		fmt.Println("n:", n, "err:", err)
		fmt.Print(name, " is ", age, " year old and has a salary of ", salary)
	*/

	/*
		// scanning from strings w/ matching format
		// example from golang.org
		var name string
		var age int
		// input "Kim is 22 years old" works but "Kim 22" and "Kim is 22" does not
		n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d\n", n, name, age)
	*/

	/*
		func Fscan(r io.Reader, a ...interface{}) (n int, err error)
		Reader is the interface that wraps the basic Read method.
			type Reader interface {
			    Read(p []byte) (n int, err error)
			}
	*/

	/*
		var animal string
		n, err := fmt.Fscan(os.Stdin, &animal)
		fmt.Println("n:", n, "err:", err)
		fmt.Print("The animal you choose is ", animal, ".\n")
	*/

	/*
		var pet1, pet2, pet3 string
		// If successful, methods on the returned file can be used for reading, *File will have a Read method, making it of type io.Reader
		file, err := os.Open("myText.txt")
		if err != nil {
			log.Fatal("error reading file:", err)
		}
		n, err := fmt.Fscan(file, &pet1, &pet2, &pet3)
		fmt.Println("n:", n, "err:", err)
		fmt.Print("Your animals are ", pet1, ", ", pet2, ", and ", pet3, ".\n")
	*/

	/*
		// math adding game
		var num1 int
		var num2 int
		var answer int
		rand.Seed(99) // deterministic, change for variability
		for {
			num1 = rand.Intn(10)
			num2 = rand.Intn(10)
		Inp:
			fmt.Print("What  is the sum of ", num1, " + ", num2, "? ")
			fmt.Scan(&answer)
			if answer == num1+num2 {
				fmt.Println("Nice work, ", num1+num2, " is correct")
			} else {
				fmt.Println("Sorry, ", num1+num2, " is not correct")
				goto Inp
			}
		}
	*/
}
