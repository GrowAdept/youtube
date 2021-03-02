// fmt package print functions
package main

func main() {

	// built in functions
	name := "Justin"
	// builtin print functions, may be removed later
	print(name + " ")
	print(name + "\n") // \n is newline escape sequence
	println(name)

	// Only "Print": prints to standard out (terminal), still formats escaped characters
	// func Print(a ...interface{}) (n int, err error)
	// func Println(a ...interface{}) (n int, err error)
	/*
		fmt.Print("Rock ")
		fmt.Print("Tree\n") // formats newline escaped characters
		fmt.Println("Sky")
		fmt.Println(44)
		fmt.Println(false)
		pet1, pet2, pet3 := "dog", "cat", "fish"
		fmt.Print("pets: ", pet1, ", ", pet2, ", ", pet3, "\n")
		n, err := fmt.Println("Soil")
		fmt.Println("n:", n, " err:", err) // prints number of bytes written and error
	*/

	// "Sprint": returns resulting string
	/*
		a := fmt.Sprint("car")
		fmt.Print(a)
		fmt.Print(a)
		b := fmt.Sprintln("car")
		fmt.Print(b)
		fmt.Print(b)
		// printing string w/ non-printable characters
		fmt.Printf("%q\n", a)
		fmt.Printf("%q\n", b)
	*/
	/*
		//"Printf: formats string with verbs"
		//General Verbs
		var e bool = true
		var f int = 10
		var g float64 = 3.13
		var h string = "car"
		var i = map[int]string{22: "Jane", 12: "Bill", 2: "Kelly"}
		fmt.Printf("value in default format: %v\n", h)
		fmt.Printf("type of e: %T\n", e)
		fmt.Printf("type of f: %T\n", f)
		fmt.Printf("type of g: %T\n", g)
		fmt.Printf("type of h: %T\n", h)
		fmt.Printf("type of i: %T\n", i)
		var myFl float64 = 83.3333
		fmt.Printf("score: %.0f%%\n", myFl) // %% is a literal percent sign, consumes no value
	*/
	/*
		// Integer Verbs
		var myInt int = 94
		fmt.Printf("base 2: %b\n", myInt)
		fmt.Printf("base 8: %o\n", myInt)
		fmt.Printf("base 10: %d\n", myInt)
		fmt.Printf("base 16: %x\n", myInt) // lower-case letters for a-f
		fmt.Printf("base 16: %X\n", myInt) // upper-case letters for A-F
		fmt.Printf("unicode format: %U\n", myInt)
	*/
	/*
		// Floating-point and complexe constituents
		var smFlt float64 = 12734.923652
		var lgFlt float64 = 8323424252352342342323.276921
		fmt.Printf("smFlt %%e: %e\n", smFlt) // scientific notation, e.g. -1.234456e+78
		fmt.Printf("smFlt %%E: %E\n", smFlt) // scientific notation, e.g. -1.234456E+78
		fmt.Printf("smFlt %%f: %f\n", smFlt) // decimal point but no exponent, e.g. 123.456
		fmt.Printf("smFlt %%F: %F\n", smFlt) // synonym for %f
		fmt.Printf("smFlt %%g: %g\n", smFlt)
		fmt.Printf("lgFlt %%g: %g\n", lgFlt)
		fmt.Printf("smFlt %%G: %G\n", smFlt)
		fmt.Printf("lgFlt %%G: %G\n", lgFlt)
	*/
	// precision with %f
	/*
		var price1, price2 float64
		price1 = 3.3
		price2 = 11.605
		fmt.Print("$", price1, "\n")
		fmt.Print("$", price2, "\n")
		fmt.Printf("$%.2f\n", price1)
		fmt.Printf("$%.2f\n", price2)
	*/
	/*
		// %f width precission for character allignment on screen
		var c = 132.3
		var d = 1.6
		fmt.Println(c)
		fmt.Println(d)
		fmt.Printf("%9.0f\n", c)
		fmt.Printf("%9.0f\n", d)
	*/
	/*
		j := "hello"
		fmt.Printf("s%%: %s\n", j)
		fmt.Printf("q%%: %q\n", j)
		fmt.Printf("x%%: %x\n", j)
		fmt.Printf("X%%: %X\n", j)
		fmt.Printf("b%%: %08b\n", j[0])
	*/
	/*
		k := []byte{'t', 'h', 'e', 'r', 'e'} // cannot be "" which would be a string and not a rune
		fmt.Printf("s%% of []byte: %s\n", k)
		fmt.Printf("q%% of []byte: %q\n", k)
		fmt.Printf("x%% of []byte: %x\n", k)
		fmt.Printf("X%% of []byte: %X\n", k)
		fmt.Printf("b%% of []byte: %08b\n", k[0])
	*/

	/*
		// %p for memory address stored in pointer
		l := "truck"
		fmt.Println(&l)
		fmt.Printf("p%%: %p", &l)
	*/
}
