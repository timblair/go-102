# Go 102 Workshop

## Concurrency

Go has been designed to simplify the complexities inherent in concurrent
programming.  When C was created, it was rare to have multiple CPUs in a single
machine, let alone having multiple cores per CPU.  Today, both are the norm,
and with the end of Moore's Law, the number of cores will only continue to
increase.

### Concurrency vs. Parallelism

There is often a confusion about concurrency and parallelism, so we'll quickly
define the two.  Concurrency is a way to structure a program by breaking it
into pieces that can be executed independently.  Parallelism is about running
those pieces simultaneously.  Concurrency is about structure; parallelism is
about execution.  Often, a concurrent problem solution can be parallelised, but
we'll focus on concurrency.

## Sections

* [Goroutines](goroutines.md)
* [Channels](channels.md)
