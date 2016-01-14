# The Basics

## Hello World

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
```

Every `.go` file must list which package it's a part of. This name is generally
the same as the repository / directory name, but for compilation to create a
binary (not just a library), you must define a file as being in `package main`,
and for that file to have a `main()` function.

The import statement is bringing in the `fmt` package from the Go standard
library, which implements functions that deal with formatted I/O.  The `main()`
function is the entry point for when the compiled binary file gets run, and it
uses the `Println` function from the `fmt` package to write to STDOUT.

## Types

Go is strongly and statically typed: all values have a specific type.

* Booleans: a boolean truth value denoted by the predeclared constants `true`
  and `false`
* Numerics: sets of integer or floating-point numbers of given sizes
* Strings: an immutable set of string values, which is a sequence of bytes
* Arrays: numbered sequence of elements of a single type; length is part of the
  type
* Slices: window view of a specific underlying array
* Maps: has table with specific types for keys and values
* Pointers: pointers to an address in memory
* Structs: a sequence of fields with a name and a type
* Channels: communication across concurrent functions by sending typed values

Booleans, numerics and strings are occasionally referred to as built-in types,
whereas arrays, slices, maps, pointers, structs and channels may be referred to
as "reference types," so-called because the "value" of a reference type is a
header value which contains a reference to the underlying data. This makes them
cheap to copy.

In the workshop today, we'll only be focussing on structs and channels, as well
as the basic types (booleans, numerics and strings).

## Basic Types

```go
b1 := true         // type is bool
b2 := false

n1 := 123          // int
n2 := 123.456      // float32/64
n3 := 1e10         // float32/64
n4 := uint8(123)   // uint
n5 := float32(123) // float32

s1 := `Raw string literal`
s2 := "Interpreted string literal"
```

## Variable Declaration

```go
var x T         // Variable x of type T with a zero value
var x T = v     // Variable x of type T with value v
var x = v       // Variable x with value v, implicit typing

x := v          // Short variable declaration (type inferred)
x, y := v1, v2  // Double declaration (similar with var)

make(T)         // make takes a type T, which must be a slice,
                // map or channel type, optionally followed by
                // a type-specific list of expressions
```

You can also use `new(T)`, which allocates zeroed storage for a new value of
type T and returns its address.

## Zero Values

Every type has a zero value. For the Rubyists, note that a boolean, numeric or
string can never be `nil`.

```go
0     // numeric
false // boolean
""    // string
nil   // pointer, channel, func,
      // interface, map, or slice
```

## Structs

A struct is a type which contains a collection of named fields. The `type`
keyword introduces a new type. It's followed by the name of the type
(`rectangle`, in the following example), the keyword `struct` to indicate that
we are defining a struct type and a list of fields inside of curly braces. Each
field has a name and a type.

```go
type rectangle struct {
	width  int
	height int
}

r1 := rectangle{1, 2}       // New rectangle with w + h
r1.width = 3                // Set width to a new value
fmt.Printf("Width = %d; Height = %d\n", r1.width, r1.height)

var r2 rectangle            // w=0, h=0 (int zero values)
r4 := rectangle{}           // w=0, h=0
r3 := rectangle{Height: 1}  // w=0, h=1
```

## Functions

A function is an independent section of code that maps zero or more input
parameters to zero or more output parameters. In Go, they are values like any
other value, and can be passed around as any other value.

```go
func f1() {}                // Simple function definition
func f2(s string, i int) {} // Function that accepts two args
func f3(s1, s2 string) {}   // Two args of the same type
func f4(s ...string) {}     // Variadic function

func f5() int {             // Return type declaration
	return 42
}

func f6() (int, string) {   // Multiple return values
	return 42, "foo"
}
```

# Exercise

> Declare a struct type to maintain information about a person.  Declare a
> function that creates new values of your type.  Call this function from main
> and display the value.

* Exercise template: [source][ts] / [playground][tp]
* Example solution: [source][ss] / [playground][sp]

[ts]: exercises/basics/template/basics.go
[tp]: http://play.golang.org/p/ta6oFzjgwn
[ss]: exercises/basics/solution/basics.go
[sp]: http://play.golang.org/p/xTcpaKL4KG
