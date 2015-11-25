# Go Workshop

## Overview

The principle goal of this workshop it to get you excited about Go! We’ll be
focussing on two key areas: object oriented development in Go, and the powerful
primitives which make concurrent programming in Go a breeze.

* [Introduction](1-the-basics.md): built-in types, variable declaration,
  function and custom types.
* [Object oriented development](2-oo-in-go.md): methods, interfaces, embedding
  and composition.
* [Concurrency](3-concurrency.md): goroutines and channels.

In each section we'll walk through a number of examples and code samples, and
will work through some coding exercises in pairs, either using a local Go
installation or the Go Playground.

## Pre-Requisites

This workshop is designed for folks who have some basic knowledge of Go, but
not necessarily any practical experience with it.  Attendees are also expected
to have practical experience with the principles behind Object Oriented
software development (in any language).

I will recap some basics at the beginning of the workshop, but you'll get more
out of it if you already have knowledge of variable declaration, functions, and
struct types. I suggest you watch the first half of my recent [introductory
talk on Go](https://vimeo.com/140410716) (the "Data & Types" section), or a
least look through the [slide
deck](https://speakerdeck.com/timblair/introduction-to-go).

## What You'll Need to Bring

As long as you have a working laptop (or a friend with one, as we'll be pairing
for the exercises), then you’re good to go.  You don’t even need Go installed
on your machine, because you can use the online [Go
Playground](http://play.golang.org/) instead for the exercises.

## Exercises

There are [templates](exercises/templates) available for each of the exercises.
Simply copy the template into a new file in an empty directory under the `src`
directory in your `$GOPATH` and fill in the blanks.  Alternatively, follow the
link to the template already set up in the Go Playground, and work from there
instead.

* [Exercise #1: The Basics](exercises/templates/01-basics.go)
* [Exercise #2: Methods](exercises/templates/02-methods.go)

Example [solutions](exercises/solutions) are also available, but don't look at
those until you've had a go yourself!
