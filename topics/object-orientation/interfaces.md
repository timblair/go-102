# Object-Oriented Programming

## Interfaces

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
type speaker interface {
	speak() string
}
```

We'll now define three new types which satisfy our `speaker` interface.  Note
that each type has methods that are specific to the type, in addition to the
required `speak` method required to satisfy the interface.

```go
type hipster struct { }
func (h hipster) speak() string { return "Amazeballs" }
func (h hipster) trimBeard() { /* ... */ }

type dog struct { }
func (d dog) speak() string { return "Woof" }
func (d dog) wagTail() { /* ... */ }

type robot struct { }
func (r robot) speak() string { return "Does not compute" }
func (r robot) becomeSentient() { /* ... */ }
```

A value of an interface type can hold any value that satisfies the methods
defined for the interface, so this means that a value of any of our
newly-defined types can be assigned to a value of type `speaker`:

```go
// We can treat a hipster as a speaker
var s1 speaker
s1 = hipster{}

// We can also create a slice of different speakers
speakers := []speaker{hipster{}, dog{}, robot{}}
for _, s := range speakers {
	fmt.Printf("%T: %s\n", s, s.speak())
}
```

Let's see a simple use of interfaces in action.  Let's define a new type which
represents an email address, including the name of the recipient:

```go
type email struct {
	name string
	address string
}

e := email{"Tim Blair", "tim@bla.ir"}
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

As a quick aside, you may notice that the interface name and defined method
above start with a capital letter.  Go doesn't have the concept of public or
private methods; instead an identifier (constant, type, function, or method) is
either _exported_ or _unexported_ from a package.  Exported identifiers are
acessible outside the package; an unexported identifier is only accessible
within the package that defines it.  You can think of them repectively as being
similar to the `public` and `package` keywords in Java.

By adding the `String` method to our type, it now satisfies the `Stringer`
interface, and the fmt package will use that method when outputting our type.

```go
type email struct {
	name string
	address string
}

func (e email) String() string {
	return fmt.Sprintf("\"%s\" <%s>", e.name, e.address)
}

e := email{"Tim Blair", "tim@bla.ir"}
fmt.Println(e)

// "Tim Blair" <tim@bla.ir>
```

### Exercise

> Define an interface which defines a method area().  Create types for square,
> rectangle and circle, and ensure they satisfy your interface.  Create a
> function that accepts a value of your interface type and outputs the area,
> and call this function for different shapes.

* Exercise template: [source][ts] / [playground][tp]
* Example solution: [source][ss] / [playground][sp]

[ts]: exercises/interfaces/template/interfaces.go
[tp]: http://play.golang.org/p/rL5tT2VTJH
[ss]: exercises/interfaces/solution/interfaces.go
[sp]: http://play.golang.org/p/bwXmMNW2ed
