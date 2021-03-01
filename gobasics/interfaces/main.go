// interfaces
package main

import (
	"fmt"
	"math"
)

type shape interface {
	perimeter() float64
	// area() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// equilateral triangle
type eqlTriangle struct {
	side float64
}

func findPerim(s shape) float64 {
	perim := s.perimeter()
	return perim
}

func (r rectangle) perimeter() float64 {
	perim := r.height + r.height + r.width + r.width
	return perim
}

func (c circle) perimeter() float64 {
	perim := 2 * math.Phi * c.radius
	return perim
}

func (t eqlTriangle) perimeter() float64 {
	perim := t.side * 3
	return perim
}

func main() {
	a := rectangle{10, 5}
	b := circle{7}
	c := eqlTriangle{3}
	fmt.Println("a.perimter returned:", a.perimeter())
	fmt.Println("b.perimter returned:", b.perimeter())
	fmt.Println("c.perimter returned:", c.perimeter())
	// fmt.Println("findPerim(a) returned:", findPerim(a))
	// fmt.Println("findPerim(b) returned:", findPerim(b))
	// fmt.Println("findPerim(c) returned:", findPerim(c))
	// fmt.Println("a.area returned:", a.area())
}

func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}
