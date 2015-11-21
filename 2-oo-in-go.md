# Go 102 Workshop

## 2. Object-Oriented Programming in Go

### OO Concepts Recap

OO is a software design philosophy, based around the concept of modelling
real-world or abstract "objects" in code.  There are four primary concepts that
are generally cited as a requirement for a language to be "object oriented":

1. *Encapsulation*: the wrapping of the data and behaviour associated with an
   object, generally achieved by creating instances of classes which expose
   public methods and properties, while restricting access to internal values.
1. *Abstraction*: 
1. *Inheritance*: is a method of code reuse, and is the ability for a class to
   "inherit" features of another class.
1. *Polymorphism* is the application of a single interface to objects of
   different types, generally in the form of subtyping (where different classes
   are related by some common superclass).

There are numerous definitions of
"an object" in terms of software, but we'll use the following simple
definition:

> An object is a data structure that has both state and behaviour.


...

Go is different to other OO languages you may have used: there are no classes,
no subclasses or class heirarchy, no explicit interfaces, no inheritence.


### Methods

Given our previous definition of an object, a custom `struct` type can contain
the state (data) of an object.  But what about the behaviour?  That behaviour
is provided by _methods_.

Let's create a `Rectangle` type, with a width and a height. From this, it's
simple to calculate the area, so let's write a function to do that.

```go
type Rectangle struct {
	Width  int
	Height int
}

func Area(r Rectangle) int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println(Area(r))    // area 12
}
```

It works as expected, but now we need to either repeat the calculation or copy
the function everywhere we want to calculate the area.  We can solve this by
turning the function into a method.

> A method is a function that is bound to a receiver.

Go does not have classes, but you can define methods on any type (that isn't a
built-in type).  The method receiver appears in its own argument list between
the func keyword and the method name.

```go
type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println(r.Area())   // area 12
}
```

Comparing the two versions, it's a very simple, yet powerful change.

```go
// function, called with Area(r)
func Area(r Rectangle) int {
	return r.Width * r.Height
}

// method, called with r.Area()
func (r Rectangle) Area() int {
	return r.Width * r.Height
}
```

It's actually worth pointing out that the `method` version is effectively just
syntactic sugar: under the hood it takes the receiver and puts it as the first
argument, just as happens in the function version.

#### Exercise #2

> Declare a new struct type to hold information about a tennis player,
> including the number of matches played and the number won.  Add a method to
> this type that calculates the win ratio for the player.  Create a new player,
> and output the win ratio for them.


### Interfaces

...

### Embedding

...

### Composition

...

> Everyone knows composition is more powerful than inheritance, Go just makes
> this non optional.
-- http://dave.cheney.net/2015/11/15/the-legacy-of-go
