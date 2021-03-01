package main

import (
	"fmt"
	"regexp"
)

func main() {

	// More Complicated queries need to use Compile and the full Regexp interface
	// instead of func MatchString or func Match

	// func MatchString(pattern string, s string) (matched bool, err error)
	matched, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println("matched:", matched, " error:", err)

	// func Match(pattern string, b []byte) (matched bool, err error)
	matched, err = regexp.Match(`\w+fo\w+`, []byte(`seafood`))
	fmt.Println("matched:", matched, " error:", err)

	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	// func Compile(expr string) (*Regexp, error)
	var re *regexp.Regexp
	re, err = regexp.Compile(`\wat`)
	fmt.Println("re:", re, "error:", err)
	str1 := `bat mat pot sat cat rat pat vat hat`
	// func (re *Regexp) MatchString(s string) bool
	fmt.Println("MatchString:", re.MatchString(str1))
	// func (re *Regexp) Find(b []byte) []byte
	fmt.Println("Find:", string(re.Find([]byte(str1))))
	// func (re *Regexp) FindAll(b []byte, n int) [][]byte
	fmt.Printf("FinadAll: %q\n", re.FindAll([]byte(str1), -1))
	// func (re *Regexp) FindIndex(b []byte) (loc []int)
	fmt.Println("FindIndex:", re.FindIndex([]byte(str1)))
	// func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
	fmt.Println("FindAllIndex:", re.FindAllIndex([]byte(str1), -1))
	// func (re *Regexp) Match(b []byte) bool
	fmt.Println("Match:", re.Match([]byte([]byte(str1))))
	// func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
	fmt.Printf("ReplaceAllLiteral: %s\n", re.ReplaceAllLiteral([]byte(str1), []byte("dog")))
	// func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
	fmt.Println("ReplaceAllLiteralString: ", re.ReplaceAllLiteralString(str1, "cat"))
	// func (re *Regexp) String() string
	fmt.Println("String:", re.String())
	var re2 *regexp.Regexp
	re2, err = regexp.Compile(`^(img-\d+)\.png$`)
	f := "img-284621831.png"
	// func (re *Regexp) ReplaceAllString(src, repl string) string
	fmt.Println("ReplaceAllString:", re2.ReplaceAllString(f, `$1`))
}
