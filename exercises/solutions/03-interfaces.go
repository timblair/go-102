// Define an interface which defines a method Area().  Create types for Square,
// Rectangle and Circle, and ensure they satisfy your interface.  Create a
// function that accepts a value of your interface type and outputs the area,
// and call this function for different shapes.
package main

// Add your imports here
import (
	"fmt"
	"math"
)

// Define an interface with a method Area().  Make sure you use a meaningful
// name and a sensible return type.
type Shaper interface {
	Area() float64
}

// Create Square, Rectangle and Circle types, and ensure they satisfy your
// interface (you'll need to use the `Pi` constant from the `math` package for
// your Circle).
type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Write a function that accepts a value of your interface and outputs the
// area.
func writeArea(s Shaper) {
	fmt.Printf("%T: %.2f\n", s, s.Area())
}

func main() {
	// Create a slice of your interface type, and populate it with a number of
	// different shapes.
	shapes := []Shaper{
		Square{5},
		Square{9},
		Rectangle{3, 7},
		Rectangle{2, 3},
		Circle{2},
		Circle{4},
	}

	// Loop through your shapes and use your function to output the area of
	// each one.
	for _, s := range shapes {
		writeArea(s)
	}
}
