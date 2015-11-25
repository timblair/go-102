// Define an interface which defines a method area().  Create types for square,
// rectangle and circle, and ensure they satisfy your interface.  Create a
// function that accepts a value of your interface type and outputs the area,
// and call this function for different shapes.
package main

// Add your imports here
import (
	"fmt"
	"math"
)

// Define an interface with a method area().  Make sure you use a meaningful
// name and a sensible return type.
type shaper interface {
	area() float64
}

// Create square, rectangle and circle types, and ensure they satisfy your
// interface (you'll need to use the `Pi` constant from the `math` package for
// your circle).
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Write a function that accepts a value of your interface and outputs the
// area.
func writeArea(s shaper) {
	fmt.Printf("%T: %.2f\n", s, s.area())
}

func main() {
	// Create a slice of your interface type, and populate it with a number of
	// different shapes.
	shapes := []shaper{
		square{5},
		square{9},
		rectangle{3, 7},
		rectangle{2, 3},
		circle{2},
		circle{4},
	}

	// Loop through your shapes and use your function to output the area of
	// each one.
	for _, s := range shapes {
		writeArea(s)
	}
}
