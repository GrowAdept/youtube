/*
sorry, forgot to put this in src/gobasics directory in the video

to build a binary in this directory run command "go build"
after building binary, in this dir run "myApp.exe" (for Windows)

to build a binary that can be run while in any directory run command "go install"
now you can run your the app in any directory with command "myApp.exe"
*/
package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}
