// scrolling to the right text, not streaming down text
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var name = "Kylee"

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func changeName(n string) {
	name = "Justin"
	fmt.Println("name:", name)
}

func ticker() {
	spacing := ""
	for {
		fmt.Println(spacing + name)
		spacing = spacing + " "
		if len(spacing) == 40 {
			spacing = ""
		}
		time.Sleep(time.Millisecond * 200)
		//CallClear()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	go ticker()
	changeName("Justin")
	fmt.Scanln()
}
