// methods (value receivers)
package main

import (
	"fmt"
)

//creates copy of type and passes it to the function, not the original
func (r rectangle) doesNotChangeRect() {
	fmt.Println("r:", r)
	r.width = r.width * 2
	r.height = r.height * 2
	fmt.Println("r:", r)
}

func area(width, height float64) float64 {
	area := width * height
	return area
}

func area2(r rectangle) float64 {
	area := r.width * r.height
	return area
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (r rectangle) perimeter() float64 {
	perimeter := r.width + r.width + r.height + r.height
	return perimeter
}

func main() {
	var a = rectangle{width: 44, height: 10}
	fmt.Println("area:", a.area())
	fmt.Println("perimeter", a.perimeter())
	// fmt.Println("area:", area(a.width, a.height))
	// fmt.Println("area:", area2(a))
	// fmt.Println("a:", a)
	// a.doesNotChangeRect()
	// fmt.Println("a:", a)
}
