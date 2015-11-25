# Go 102 Workshop

## Object-Oriented Programming in Go

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

## Sections

* [Methods](methods.md)
* [Interfaces](interfaces.md)
* [Embedding](embedding.md)
* [Composition](composition.md)
