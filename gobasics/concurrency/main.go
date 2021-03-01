package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var name = "Kylee"

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
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	go ticker()
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		// non windows users
		// text = strings.Replace(text, "\n", "", -1)

		// Windows users
		text = strings.Replace(text, "\r\n", "", -1)
		name = text

	}

	// var i int
	// fmt.Scanf("%s", &i)
	// fmt.Println(string(i))
	changeName("Justin")
	fmt.Scanln()
}
