# Go 102 Workshop

## Object-Oriented Programming in Go

### Interfaces

In Go, interfaces provide polymorphism, just like in any other OO language, by
declaring the behaviour of a type, and an interface is defined by the set of
methods it declares.  We use an interface to specify the behaviour of a given
object.

Like most uses of OO languages, the relevant behaviour is defined in a concrete
type via methods, but Go uses interfaces differently than most other OO
languages in two key ways: there is no `implements` keyword, and interfaces are
satisfied _implicitly_.

Let's define a simple interface which defines the method for allowing the
object in question to "speak." In Go, by convention, one-method interfaces are
named by the method name plus an -er suffix.

```go
type Speaker interface {
	Speak() string
}
```

We'll now define three new types which satisfy our `Speaker` interface.  Note
that each type has methods that are specific to the type, in addition to the
required `Speak` method required to satisfy the interface.

```go
type Hipster struct { }
func (h Hipster) Speak() string { return "Amazeballs" }
func (h Hipster) TrimBeard() { /* ... */ }

type Dog struct { }
func (d Dog) Speak() string { return "Woof" }
func (d Dog) WagTail() { /* ... */ }

type Robot struct { }
func (r Robot) Speak() string { return "Does not compute" }
func (r Robot) BecomeSentient() { /* ... */ }
```

A value of an interface type can hold any value that satisfies the methods
defined for the interface, so this means that a value of any of our
newly-defined types can be assigned to a value of type `Speaker`:

```go
// We can treat a Hipster as a Speaker
var s1 Speaker
s1 = Hipster{}

// We can also create a slice of different Speakers
speakers := []Speaker{Hipster{}, Dog{}, Robot{}}
for _, s := range speakers {
	fmt.Printf("%T: %s\n", s, s.Speak())
}
```

Let's see a simple use of interfaces in action.  Let's define a new type which
represents an email address, including the name of the recipient:

```go
type Email struct {
	name string
	address string
}

e := Email{"Tim Blair", "tim@bla.ir"}
fmt.Println(e)

// {Tim Blair tim@bla.ir}
```

We can pass our value to `fmt.Println` and it will happily output the
information from the struct, but maybe it's not quite in the format we'd like.
If we want to change this output format, we can make use of the `Stringer`
interface defined in the `fmt` package:

```go
// The Stringer interface found in fmt package
type Stringer interface {
	String() string
}
```

By adding the `String` method to our type, it now satisfies the `Stringer`
interface, and the fmt package will use that method when outputting our type.

```go
type Email struct {
	name string
	address string
}

func (e Email) String() string {
	return fmt.Sprintf("\"%s\" <%s>", e.name, e.address)
}

e := Email{"Tim Blair", "tim@bla.ir"}
fmt.Println(e)

// "Tim Blair" <tim@bla.ir>
```

#### Exercise

> Define an interface which defines a method area().  Create types for square,
> rectangle and circle, and ensure they satisfy your interface.  Create a
> function that accepts a value of your interface type and outputs the area,
> and call this function for different shapes.
