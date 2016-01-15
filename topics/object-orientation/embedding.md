# Object-Oriented Programming

## Embedding

Go's type system does not typical OO inheritance; instead it supports
composition with the ability to "borrow" functionality by embedding named types
as anonymous fields of another struct type through a process called _type
embedding_.

Let's define two types that are shapes inside a 3D drawing package:

```go
type sphere struct {
	x, y, z, radius int
}

type cube struct {
	x, y, z, length int
}
```

Looking at these two types, and considering what new types we might want to
create, we can see that the location of the shape is common to all shapes, so
we're going to get a lot of repetition.

We can factor out the common parts (the location) by creating a `point` type,
but this can make accessing the fields cumbersome and verbose, especially if we
then used these types as part of further types.

```go
type point struct {
	x, y, z int
}

type sphere struct {
	point point
	radius int
}

type cube struct {
	point point
	length int
}

var s sphere
s.point.x = 5
s.point.y = 6
s.point.z = 7
s.radius = 3
```

Go allows us to declare a field with a type but no name.  These fields are
called _anonymous fields_, and are said to be _embedded_ within the type.  For
example, we can embed `point` within `sphere` and `cube`.


```go
type point struct {
	x, y, z int
}

type sphere struct {
	point
	radius int
}

type cube struct {
	point
	length int
}
```

Go automatically "promotes" the embedded types so the fields within those types
are accessible from the top level.

```go
var s sphere
s.x = 5       // equivalent to s.point.x = 5
s.y = 6       // equivalent to s.point.y = 6
s.z = 7       // equivalent to s.point.z = 7
s.radius = 3
```

Although embedded fields are known as _anonymous_, that's not strictly true:
they do have names, which are the same as the type name, and using these names
is optional in dot expressions.

There is no shorthand for defining struct literals with embedded types, so we
have to declare each embedded type explicitly.

```go
s1 := sphere{point{1, 2, 3}, 5}

s2 := sphere{
	point: point{
		1,
		2,
		3,
	},
	radius: 5,   // required trailing comma
}
```

The inner type promotion works not just for the fields of embedded types, but
also for methods on the inner type. This is the primary mechanism through which
complex object behaviours are composed from simpler ones in Go.

```go
type robot struct { }
func (r robot) talk() { fmt.Println("Bzzzzzbt") }

type robby struct {
	robot
}

robby := robby{}
robby.talk()   // Bzzzzzbt
```

The outer type can override the inner type's behaviour.

```go
type robot struct { }
func (r robot) talk() { fmt.Println("Bzzzzzbt") }

type robby struct {
	robot
}
func (r robby) talk() { fmt.Println("Again?") }

robby := robby{}
robby.talk()   // Again?
```

And methods promoted from an inner type can also allow the outer type so
satisfy an interface.

```go
type talker interface { talk() }

type robot struct{}
func (r robot) talk() { fmt.Println("Bzzzzzbt") }

type robby struct { robot }

func talk(t talker) {
	t.talk()
}

talk(robby{})
```

### Exercise

> Create a user type, and an admin type that embeds a user. Create a Notifier
> interface, and make your user type satisfy that interface. Write a function
> that accepts a value of the interface type, and ensure it works correctly
> when passed a value of your admin type.

* Exercise template: [source][ts] / [playground][tp]
* Example solution: [source][ss] / [playground][sp]

[ts]: exercises/embedding/template/embedding.go
[tp]: http://play.golang.org/p/5qrrcfHdiZ
[ss]: exercises/embedding/solution/embedding.go
[sp]: http://play.golang.org/p/BU033x1m6s
