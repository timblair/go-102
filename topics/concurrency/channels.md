# Concurrency

## Channels

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
material](https://github.com/ardanlabs/gotraining/tree/master/topics/go/concurrency/channels)).

Unbuffered channel:

![Unbuffered](https://raw.githubusercontent.com/ardanlabs/gotraining/master/topics/go/concurrency/channels/unbuffered.png)

An unbuffered channel is also commonly referred to as a `synchronous` channel.

Buffered channel:

![Buffered](https://raw.githubusercontent.com/ardanlabs/gotraining/master/topics/go/concurrency/channels/buffered.png)

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
	court := make(chan struct{}) // An unbuffered channel.
	wg.Add(2) // Add two to the WG, one for each player.

	// Launch two players.
	go player("Serena", court)
	go player("Venus", court)

	court <- struct{}{} // Serve the "ball."
	wg.Wait()  // Wait for the game to finish.
}
```

Now we'll define the function that represents the player.  We'll start simple:
we create an infinite loop that waits for the ball (receives on the `court`
channel), and then hits it back (sends to the channel).

```go
func player(name string, court chan struct{}) {
	for {
		ball := <-court
		fmt.Println(name, "hit the ball")
		court <- ball // Hit the ball back.
	}
}
```

We're going to use the notification mechanism we described previously about
closing the channel to signify the end of the point: if a player tries to
recieve the ball, and the channel is closed, then the other player missed the
ball.  We change our receive to the two-value format; if `ok` is `false`, then
the channel is closed, and our player won the point!


```go
func player(name string, court chan struct{}) {
	for {
		ball, ok := <-court
		if !ok { // If the channel was closed we won.
			fmt.Println(name, "won!")
			return
		}

		fmt.Println(name, "hit the ball")
		court <- ball // Hit the ball back.
	}
}
```


With that, no one would ever miss the ball, so the point would go on
forever!  Let's add something to the `player()` function to randomly determine
if the player misses the ball.  If they do, we'll close the `court` channel to
signify that the point is over.

```go
func player(name string, court chan struct{}) {
	for {
		// Receive step excluded...

		if rand.Intn(10) == 0 {  // Decide if we missed the ball.
			fmt.Println(name, "missed the ball")
			close(court) // Close the channel to signal we lost.
			return
		}

		fmt.Println(name, "hit the ball")
		court <- ball // Hit the ball back.
	}
}
```

We're almost done.  The players can hit the ball back, may (or may not) miss
the ball, and can recognise when the other player has missed.

The only thing missing is that fact that our WaitGroup is still blocked,
because we've never decreased the count from it, so our `main()` routine will
block forever.  It's like the players know what's going on, but the umpire has
fallen asleep in his chair.

That's easily solved by defering a call to `wg.Done()` within the player
function, so each time a player goroutine finishes (i.e. we hit a `return`
statement), that semaphore is decreased.

```go
func player(name string, court chan struct{}) {
	defer wg.Done()

	for {
        // ...
	}
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

A complete, runnable version of the above example can be found on the [Go
Playground](http://play.golang.org/p/Q4z-QyL4D8).

### Exercise

> Let's simulate a track relay race.  Create a channel representing a track,
> and a function representing a runner.  Pass a baton between runners over the
> channel, and end the race when the fourth runner receives the baton.

* Exercise template: [source][ts] / [playground][tp]
* Example solution: [source][ss] / [playground][sp]

[ts]: exercises/channels/template/channels.go
[tp]: http://play.golang.org/p/H4F9aLKQVA
[ss]: exercises/channels/solution/channels.go
[sp]: http://play.golang.org/p/eNglulpfz2
