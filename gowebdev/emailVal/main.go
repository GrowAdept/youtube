package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

// func MustCompile(str string) *Regexp
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func main() {

	emails := []string{
		"!#$%&'*+-/=?^_`{|}~0123456789@mail.com", // allowed characters in local part !#$%&'*+-/=?^_`{|}~
		"janedoe@email.com",                      // uppercase and lowercase Latin letters A to Z and a to z allowd in local part
		"JaneDoe@email.com",                      // // uppercase and lowercase Latin letters A to Z and a to z allowd in local part
		"Jane.Doe@email.com",                     // . allowed in local part as long as it's not the first, last, or there are consectutives
		"Jane-Doe@email.com",                     // allowed characters in local part !#$%&'*+-/=?^_`{|}~
		"a@email.com",
		// "abcdedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzazzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz@mail.com", // local longer than 64 characters
		"JaneDoeemail.com",           // must have @
		"Jane@Doe@email.com",         // cannot have more than one @
		"Jane:Doe@email.com",         // (),:;<>@[\] not allowed outsided of qoutes
		"janedoe@email.-com",         // - cannot be first or last character in domain
		"janedoe@email.com-",         // - cannot be first or last character in domain
		"janedoe@email.)com",         // domain may only contain a-z, A-Z, 0-9, or hyphen
		"janedoe@mysite_email.!com",  // domain may only contain a-z, A-Z, 0-9, or hyphen
		"janedoe@mysite+account.com", // domain may only contain a-z, A-Z, 0-9, or hyphen
	}

	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', tabwriter.Debug)

	fmt.Fprint(w, "Pass Regex\tPass Length\tEmail Address\t\n")
	for _, v := range emails {
		// func (re *Regexp) MatchString(s string) bool
		rg := emailRegex.MatchString(v)
		ln := len(v) > 3 && len(v) < 254
		// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
		fmt.Fprintln(w, rg, "\t", ln, "\t", v, "\t")
	}
	w.Flush()

	addr := os.Getenv("ToEmailAddr")
	// addr := "johndoe@nothisisnotadomain.com"
	// func Split(s, sep string) []string
	i := strings.Index(addr, "@")
	fmt.Println("i:", i)
	host := addr[i+1:]
	fmt.Println(host)
	// func LookupMX(name string) ([]*MX, error)
	_, err := net.LookupMX(host)
	if err != nil {
		// handle err
	}
	fmt.Println(err)

}

/*
^ start of string
[] character class show range of characters
() capture group
{} quantifier defines quantities of our data

regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")^
^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@
[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9]
[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])

janedoe@gmail.com
john.smith@yahoo.com
abcdedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.!#$%&'*+/=?^_`{|}~-@email.com
*/
