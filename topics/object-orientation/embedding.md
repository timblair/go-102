# Go 102 Workshop

## Object-Oriented Programming in Go

### Embedding

Go's type system does not typical OO inheritance; instead it supports
composition with the ability to "borrow" functionality by embedding named types
as anonymous fields of another struct type through a process called _type
embedding_.

Let's define two types that are shapes inside a 3D drawing package:

```go
type Sphere struct {
	X, Y, Z, Radius int
}

type Cube struct {
	X, Y, Z, Length int
}
```

Looking at these two types, and considering what new types we might want to
create, we can see that the location of the shape is common to all shapes, so
we're going to get a lot of repetition.

We can factor out the common parts (the location) by creating a `Point` type,
but this can make accessing the fields cumbersome and verbose, especially if we
then used these types as part of further types.

```go
type Point struct {
	X, Y, Z int
}

type Sphere struct {
	Point Point
	Radius int
}

type Cube struct {
	Point Point
	Length int
}

var s Sphere
s.Location.X = 5
s.Location.Y = 5
s.Radius = 3
```

Go allows us to declare a field with a type but no name.  These fields are
called _anonymous fields_, and are said to be _embedded_ within the type.  For
example, we can embed `Point` within `Sphere` and `Cube`.


```go
type Point struct {
	X, Y, Z int
}

type Sphere struct {
	Point
	Radius int
}

type Cube struct {
	Point
	Length int
}
```

Go automatically "promotes" the embedded types so the fields within those types
are accessible from the top level.

```go
var s Sphere
s.X = 5       // equivalent to s.Point.X = 5
s.Y = 5       // equivalent to s.Point.Y = 5
s.Z = 5       // equivalent to s.Point.Z = 5
s.Radius = 3
```

Although embedded fields are known as _anonymous_, that's not strictly true:
they do have names, which are the same as the type name, and using these names
is optional in dot expressions.

There is no shorthand for defining struct literals with embedded types, so we
have to declare each embedded type explicitly.

```go
s1 := Sphere{Point{1, 2, 3}, 5}

s2 := Sphere{
	Point: Point{
		1,
		2,
		3,
	},
	Radius: 5,   // required trailing comma
}
```

The inner type promotion works not just for the fields of embedded types, but
also for methods on the inner type. This is the primary mechanism through which
complex object behaviours are composed from simpler ones in Go.

```go
type Robot struct { }
func (r Robot) Talk() { fmt.Println("Bzzzzzbt") }

type Robby struct {
	Robot
}

robby := Robby{}
robby.Talk()   // Bzzzzzbt
```

The outer type can override the inner type's behaviour.

```go
type Robot struct { }
func (r Robot) Talk() { fmt.Println("Bzzzzzbt") }

type Robby struct {
	Robot
}
func (r Robot) Talk() { fmt.Println("Again?") }

robby := Robby{}
robby.Talk()   // Again?
```

And methods promoted from an inner type can also allow the outer type so
satisfy an interface.

```go
type Talker interface { Talk() }

type Robot struct{}
func (r Robot) Talk() { fmt.Println("Bzzzzzbt") }

type Robby struct { Robot }

func talk(t Talker) {
	t.Talk()
}

talk(Robby{})
```

#### Exercise

> Create a user type, and an admin type that embeds a user. Create a Notifier
> interface, and make your user type satisfy that interface. Write a function
> that accepts a value of the interface type, and ensure it works correctly
> when passed a value of your admin type.
