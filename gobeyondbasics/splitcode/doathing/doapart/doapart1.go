package doapart

import "fmt"

func init() {
	fmt.Println("2 pkg: doapart  file: doapart.go  msg: this is an init function")
}

// needs to be capitalized to run outside of package
func PrintAnotherProv(prov proverb2) {
	fmt.Println("9 pkg: doapart  file: doapart1.go msg:", prov)
}
