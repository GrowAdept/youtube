package doapart

import "fmt"

func init() {
	fmt.Println("2 pkg: doapart  file: doapart.go  msg: this is an init function")
}

// needs to be capitalized to be exported outside of package
var P proverb2 = "A little copying is better than a little dependency."
