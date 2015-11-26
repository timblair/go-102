# Go 102 Workshop

## Object-Oriented Programming in Go

### Methods

Given our previous definition of an object, a custom `struct` type can contain
the state (data) of an object.  But what about the behaviour?  That behaviour
is provided by _methods_.

Let's create a `rectangle` type, with a width and a height. From this, it's
simple to calculate the area, so let's write a function to do that.

```go
type rectangle struct {
	width  int
	height int
}

func area(r rectangle) int {
	return r.width * r.height
}

func main() {
	r := rectangle{3, 4}
	fmt.Println(area(r))    // area 12
}
```

It works as expected, but now we need to either repeat the calculation or copy
the function everywhere we want to calculate the area.  We can solve this by
turning the function into a method.

> A method is a function that is bound to a receiver.

Go does not have classes, but you can define methods on any type.  The method
receiver appears in its own argument list between the func keyword and the
method name.

```go
type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	r := rectangle{3, 4}
	fmt.Println(r.area())   // area 12
}
```

Comparing the two versions, it's a very simple, yet powerful change.

```go
// function, called with Area(r)
func area(r rectangle) int {
	return r.width * r.height
}

// method, called with r.Area()
func (r rectangle) area() int {
	return r.width * r.height
}
```

It's actually worth pointing out that the `method` version is effectively just
syntactic sugar: under the hood it takes the receiver and puts it as the first
argument, just as happens in the function version.

#### Exercise

> Declare a new struct type to hold information about a tennis player,
> including the number of matches played and the number won.  Add a method to
> this type that calculates the win ratio for the player.  Create a new player,
> and output the win ratio for them.

* Exercise template: [source][ts] / [playground][tp]
* Example solution: [source][ss] / [playground][sp]

[ts]: exercises/methods/template/methods.go
[tp]: http://play.golang.org/p/jnBw-jtE3n
[ss]: exercises/methods/solution/methods.go
[sp]: http://play.golang.org/p/IxxPLmfdE9
