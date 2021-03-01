// Errors
package main

import (
    "fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("doesnotexist.txt")
	fmt.Println("err:", err)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("Contents of file:\n", string(content))
}