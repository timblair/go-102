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

#### Exercise #6

...


### Channels

When using the traditional threading models commonly used in Java, C++ and
Python, for example, communication between threads is usually performed through
the use of access to share data structures which are protected by locks to
avoid race conditions and inconsistencies in data.

Go simplifies this error-prone method by encouraging an approach in which
shared _values_ are passed around on _channels_.  Only one goroutine has access
to a given value at any given time, so concurrent access to the value is
avoided, by design.

Go's approach to communication through the use of channels can be neatly
summarised in the following slogan:

> "Do not communicate by sharing memory; instead, share memory by
> communicating." — https://golang.org/doc/effective_go.html#concurrency

There are cases where a lock-based approach is the better one to take
(reference counting, or for concurrent access to a map, for example) but, in
general, using channels makes it easier to write clear, correct programs.

Channel values are defined using the `make()`, and we send or receive
communications using the `<-` operator: between the channel and value for
sends, and preceding the channel for receives.  Channels are typed like any
other value in Go.  Trying to send a `string` value on a channel defined with
`int` results in a compile-time error.

```go
ch := make(chan int)
ch <- 42             // Send a value to the channel
v := <-ch            // Receive a value from ch
```

If you tried to run the previous example, it would compile fine, but when you
run it the program would panic (effectively, crash) with the error message "all
goroutines are asleep - deadlock!"  This is because when we created the channel,
we defined it as _unbuffered_.

Channels can be buffered or unbuffered.  A send on an unbuffered channel blocks
until another goroutine executes a receive on the same channel; at this point
the value is transmitted and both goroutines continue.  A buffered channel
has a queue of elements, the size of which is defined at the point the channel
is created.

The second argument to the `make()` call defines the maximum size of the buffer
available.  If it is excluded or zero, an unbuffered channel is created.

```go
c1 := make(chan int)       // Unbuffered
c2 := make(chan int, 0)    // Unbuffered
c3 := make(chan int, 100)  // Buffered
```

The differences behind the operation of buffered and unbuffered channels is
neatly illustrated by the following diagrams (courtesy of [Ardan Labs'
fantastic Go Training course
material](https://github.com/ardanlabs/gotraining/tree/master/topics/channels)).

Unbuffered channel:

![Unbuffered](https://raw.githubusercontent.com/ardanlabs/gotraining/master/topics/channels/unbuffered.png)

An unbuffered channel is also commonly referred to as a `synchronous` channel.

Buffered channel:

![Buffered](https://raw.githubusercontent.com/ardanlabs/gotraining/master/topics/channels/buffered.png)

It's tempting to use a buffered channel within a single goroutine as a simple
queue, but it's very easy for your program to become deadlocked without another
goroutine to receive on the channel, so avoid this behaviour.

It's also important to point out that once a channel reaches its maximum
capacity, further sends to the channel will block, just like an unbuffered
channel, so it's often best to start with an unbuffered channel to ensure that
you cater for this case.  In general, try to avoid buffered queues unless you
already know the upper bound for number of values that will be sent on a
channel, and a buffer with an arbitrary size should be treated as a bad code
smell.

Based on the knowledge of the differences between buffered and unbuffered
channels, we can alter our previous example to use a buffered channel so it no
longer blocks.

```go
c2 := make(chan int, 1)  // A buffered channel
c2 <- 42
v := <-c2
```

We can use an unbuffered channel as a signalling mechanism.  This use-case is
an alternative to the use of `sync.WaitGroup` we saw when discussing
goroutines.  The use of a channel here actually simplifies the code.

The `main()` function creates an unbuffered channel, and then passes that to
the `doWork()` function which is launched as a goroutine.  We then read from
the channel straight away, which blocks until `doWork()` has finised what it's
doing and has published to the channel, at which point the program exits.

```go
func doWork(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done.")
	done <- true
}

func main() {
	done := make(chan bool)
	go doWork(done)
	<-done
}
```

Here's an example of a use-case for a buffered channel.  We're implementing a
parallel search where we perform three searches, and return the result from the
fasted query to respond.  If we used an unbuffered channel, the second two
goroutines would have blocked forever trying to send to a channel that will
never be read from.  This is known as a _goroutine leak_, as goroutines aren't
garbage collected like variables.

```go
func search(q string, server string) string { /* ... */ }

func parallelSearch(q string, servers []string) string {
	res := make(chan string, 3)
	for _, s := range servers {
		go func(x string) { res <- search(q, x) }(s)
	}
	return <-res
}

servers := []string{"s1", "s2", "s3"}
fmt.Println(parallelSearch("foo", servers))
```

You can close a channel for writing using the `close()` built-in function.
Closed channels can also provide a notification mechanism by using a two-value
version of a receive, where the second value will be `true` is the receive was
successful, and `false` if the channel has been drained and is now closed.

```go
ch := make(chan int)

close(ch)        // Close the channel for writing
v, ok := <-ch    // Receive, testing for closure
if !ok {
	  fmt.Println("Channel closed!")
}
```

The previous pattern is often used for iterating over values in a channel and,
because it is common, we can instead use a `range` query to perform the same.

```go
ch := make(chan int)

// Read from channel until it is closed
for i := range ch {
	fmt.Println(i)
}
```

Let's model something slightly more complicated: a tennis game.  We'll start
with the `main()` function that kicks things off.  We create a court, set up
two players on the court, serve the ball, and then wait for someone to win the
point.

```go
var wg sync.WaitGroup

func main() {
	court := make(chan struct{}) // Create an unbuffered channel.
	wg.Add(2) // Add a count of two, one for each goroutine.

	// Launch two players.
	go player("Serena", court)
	go player("Venus", court)

	court <- struct{}{} // Start the set.
	wg.Wait()  // Wait for the game to finish.
}
```

Now we'll define the function that represents the player.  First, we wait to
receive the ball; if the `court` channel has been closed, we use this to
signify that the other player missed the ball, so this player is the winner.
We then send the ball back on the channel to signify a successful return.

```go
func player(name string, court chan struct{}) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok { // If the channel was closed we won.
			fmt.Println(name, "won!")
			return // Trigger the deferred wg.Done()
		}

		fmt.Println(name, "hit the ball")
		court <- ball // Hit the ball back.
	}
}
```

With what we have, no one would ever miss the ball, so the point would go on
forever!  Let's add something to the `player()` function to randomly determine
if the player misses the ball.  If they do, we'll close the `court` channel to
signify that the point is over.

```go
// Pick a random number and see if we miss the ball.
if rand.Intn(10) == 0 {
	fmt.Println(name, "missed the ball")
	close(court) // Close the channel to signal we lost.
	return       // Trigger the deferred wg.Done()
}
```

Putting it all together, we can run our point.

```
$ go run tennis.go
Venus hit the ball
Serena hit the ball
Venus hit the ball
Serena hit the ball
Venus hit the ball
Serena hit the ball
Venus missed the ball
Serena won!
```

#### Exercise #7

...
