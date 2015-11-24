# Go 102 Workshop

## 3. Concurrency

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

### Goroutines

A goroutine is simply a function or method running independently (but in the
same address space as other goroutines).  The Go scheduler multiplexes
goroutines against one or more OS threads, which are in turn managed by the OS
to assign work to individual CPU cores.  It's perfectly normal to create
thousands of goroutines in a single process; they're not free, but they're very
cheap.

Let's create a function to simulate some form of work, and we'll call that
multiple times.

```go
func doWork(i int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(i)
}

func main() {
	for i := 1; i <= 5; i++ {
		doWork(i)
	}
}
```

If we run that script, it does exactly what we'd expect: outputs each number
one after the other, in order, and takes approximately 5 times the time it
would take to run once.

```
$ time go run g1.go
1
2
3
4
5
go run g1.go ... 2.757 total
```

It's simple to launch any function or method call as a goroutine: simply add
the `go` keyword in front of the call.

```go
go someFunc() // Concurrency!
```

Here's out previous program, and we can see that we've added the `go` keyword
in the loop, so we're launching each of our invocations of `doWork()` in its
own goroutine.

```go
func doWork(i int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(i)
}

func main() {
	for i := 1; i <= 5; i++ {
		go doWork(i)   // Concurrency!
	}
}
```

So let's run the script again!

```
$ time go run g2.go
go run g2.go ... 0.247 total
```

Uh oh, where's all our output gone?  Here's an important thing about goroutines:
when you start a goroutine, the calling code doesn't wait for the goroutine to
finish, but just continues running through the rest of the code.  The `main()`
function is actually a goroutine itself, so once that finishes, the program
ends.

Let's introduce one method of solving this problem: the wait group.  A wait
group is a counting semaphore.  We use `wg.Add(n)` to increment the count, and
`wg.Done()` to decrement it.  A call to `wg.Wait()` will block the calling
goroutine until the wait group's value is zero.

```go
import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// Do work...
		wg.Done()
	}()

	wg.Wait()
}
```

We've also introduced two new concepts here: anonymous functions (a function
that is not bound to an identifier), and closures (functions that can access
variables local to the scope it was created in, e.g. the `wg` value).

So let's apply that concept to our previous example.

```go
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()
			doWork(x)
		}(i)
	}

	wg.Wait()
}
```

We're adding to the WaitGroup on each loop, then calling an anonymous function
with the current value of `i` (if we just used `i` directly, then due to the
closure the value passed to our `doWork()` function would be the value of `i`
at the point the function was called, which may have changed from when the
anonymous function was defined).

We've also added a call to `defer`.  This statement pushes a function call on
to a list, and the list of saved calls is executed after the surrounding
function returns.  It is commonly used to simplify functions that perform
various clean-up operations, such as closing a file handle, or to decrement a
WaitGroup like in our example.  Deferred statements run even if the function
results in a panic.

Now we've done all that, let's run our program again.

```
$ time go run g3.go
4
1
5
2
3
go run g3.go ... 0.752 total
```

We can see that each goroutine is now running concurrently with the others, and
the total time taken is approximately that for a single call to `doWork()`.

#### Exercise #5

...


### Channels

...
