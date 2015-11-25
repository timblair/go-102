# Go 102 Workshop

## 2. Object-Oriented Programming in Go

From the [Go FAQ](https://golang.org/doc/faq#Is_Go_an_object-oriented_language):

> Is Go an object-oriented language?  Yes and no.

Go has no classes or type hierarchy, no inheritance or subclassing, and
no explicit interfaces. But Go has types and methods and allows an
object-oriented style of programming.

OO is a software design philosophy, based around the concept of modelling
real-world or abstract "objects" in code.  There are four primary principles
that are generally cited as a requirement for a language to be "object
oriented":

1. *Abstraction* is the concept of exposing only the essential characteristics
   and behaviour of an object to a collaborator, so as to provide a simple
   external interface.

1. *Encapsulation* is complementary to abstraction.  Whereas abstraction
   focuses on the observable behaviour of an object, encapsulation focuses on
   the implementation of that object which gives rise to the behaviour, through
   the wrapping of data and behaviour into a single unit (usually a class).

1. *Inheritance* is a method of code reuse, and is the ability for a class to
   "inherit" features of another class through subclassing.

1. *Polymorphism* is the application of a single interface to objects of
   different types, allowing those objects to be treated similarly based on
   their behaviour rather than their type.

Although these principles are taken by many to be absolutely essential for any
language which claims to support OOP, we're going to scale things back.  At the
simplest level, OOP is about objects.  There are numerous definitions of "an
object" in terms of software, but we'll use the following simple definition:

> An object is a data structure that has both state and behaviour.

Go's type system and the ability to create methods on types gives us the
ability to create these "objects," so let's see how it's done.


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

Go does not have classes, but you can define methods on any type.  The method
receiver appears in its own argument list between the func keyword and the
method name.

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

#### Exercise #3

> Define an interface which defines a method area().  Create types for square,
> rectangle and circle, and ensure they satisfy your interface.  Create a
> function that accepts a value of your interface type and outputs the area,
> and call this function for different shapes.


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

#### Exercise #4

> Create a user type, and an admin type that embeds a user. Create a Notifier
> interface, and make your user type satisfy that interface. Write a function
> that accepts a value of the interface type, and ensure it works correctly
> when passed a value of your admin type.


### Composition

We've now introduced the building blocks for Go's powerful composition
techniques: custom types, methods, interfaces, and struct embedding.

> "Everyone knows composition is more powerful than inheritance, Go just makes
> this non optional." — Dave Cheney: http://bit.ly/dctlg

Composition is more than just type embedding; it's about declaring and
implementing discrete behaviour, creating types that have a single purpose,
and then using these blocks to build higher level behaviours by composing the
discrete behaviours.

Let's put all these techniques together.  First we'll define a type
representing a location, and two interfaces detailing behaviour.

```go
type Point struct {
	X, Y int
}

type Mover interface {
	MoveTo(p Point)
}

type Firer interface {
	Fire()
}
```

Now we'll define a new type for a vehicle, embedding the location.  We'll also
make this type satisfy one of our interfaces.

```go
type Vehicle struct {
	Point
	Passengers int
}

func (v *Vehicle) MoveTo(p Point) {
	v.Point = p
}
```

Let's also have a weapon, again satisfying the appropriate interface.

```go
type Weapon struct {
	Loaded bool
}

func (w *Weapon) Fire() {
	w.Loaded = false
}
```

With a bit more struct embedding, we can compose our `Vehicle` and `Weapon`
types to create a `Tank`.

```go
type Tank struct {
	Vehicle
	Weapon
}
```

Here we'll introduce a familiar concept with a slight twist: it's type
embedding, but with an interface.  It works in exactly the same way as with
a struct.

```go
type MoverFirer interface {
	Mover
	Firer
}

func moveAndFire(mf MoverFirer, p Point) {
	mf.MoveTo(p)
	mf.Fire()
}
```

And finally, let's wrap everything together.  We'll create a new `Tank`, and
then use our `moveAndFire` function to do just that.

```go
func main() {
	t := &Tank{
		Vehicle{Point{5, 6}, 6},
		Weapon{true},
	}

	moveAndFire(t, Point{10, 20})

	fmt.Printf("Location: %v; Passengers: %d; Loaded: %t\n",
		t.Point, t.Passengers, t.Loaded)

	// Location: {10 20}; Passengers: 6; Loaded: false
}
```

#### Exercise #5

...
