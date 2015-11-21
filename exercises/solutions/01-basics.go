// Declare a struct type to maintain information about a person.  Declare a
// function that creates new values of your type.  Call this function from main
// and display the value.
package main

// Add your imports here
import "fmt"

// Declare a struct type `person` to maintain information about a person.
type person struct {
	name string
	age  int
}

// Declare a function that creates new values of your `person` type.
func NewPerson(name string, age int) person {
	return person{name, age}
}

func main() {
	// Use you function to create a new value of type `person`.
	p := NewPerson("Bobby", 22)

	// Output the value of your person.
	fmt.Println(p)
}
