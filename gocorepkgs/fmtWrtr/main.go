// fmt package io.Writer functions
package main

import (
	"fmt"
	"os"
)

func main() {

	// printing to standard out (terminal)
	proverb := "Design the architecture, name the components, document the details."
	// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	fmt.Fprint(os.Stdout, proverb)

	/*
		// printing to file
		proverb1 := "Design the architecture, name the components, document the details.\n"
		// func OpenFile(name string, flag int, perm FileMode) (*File, error)
		f, err := os.OpenFile("myFile.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
		fmt.Println("err:", err)
		defer f.Close()
		fmt.Fprint(f, proverb1)
	*/

	/*
		// proverb2 := "Documentation is for users.\n"
		l, _ := os.OpenFile("logFile.txt", os.O_RDWR|os.O_APPEND, 0666)
		log.SetOutput(l)
		// _, err = fmt.Fprint(f, proverb2)
		err = errors.New("we have a problem")
		if err != nil {
			fmt.Println("there was an error")
			log.Fatalln(err)
		}
		defer l.Close()
	*/
}
