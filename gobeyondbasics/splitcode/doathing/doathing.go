package doathing

import "fmt"

func init() {
	fmt.Println("1 pkg: doathing file: doathing.go msg: this is an init function")
}

var Proverb2 = "interface{} says nothing."
var proverb3 = "Reflection is never clear."

func PrintProverb(prov string) {
	fmt.Println("8 pkg: doathing file: doathing.go msg:", prov)
}
