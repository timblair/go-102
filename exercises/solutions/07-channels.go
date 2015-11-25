// Let's simulate a track relay race.  Create a channel representing a track,
// and a function representing a runner.  Pass a baton between runners over the
// channel, and end the race when the fourth runner receives the baton.
package main

// Add your imports here.
import (
	"fmt"
	"sync"
)

// Create a new waitgroup.
var wg sync.WaitGroup

func main() {
	// Add something to the waitgroup so main() can wait for the race to
	// finish.
	wg.Add(1)

	// Create an unbuffered channel to act as the track.  In terms of the
	// channel type, think about how you're going to count the runners.
	track := make(chan int)

	// Put the first runner at his starting position by launching the
	// `runner()` function as a goroutine.
	go runner(track)

	// Give the runner the baton by sending it on the channel.
	track <- 1

	// Wait until the race is over.
	wg.Wait()

	fmt.Println("The race is over")
}

// Create the function representing the runner.  They'll need to be given the
// track channel so we can pass the baton between runners.
func runner(track chan int) {
	fmt.Println("New runner ready to receive the baton")

	// Set the runner ready to receive the baton by receiving on the channel.
	baton := <-track

	fmt.Println("Runner is running leg", baton)

	// If this runner is the fourth runner then the race is over.
	if baton == 4 {
		// Signify the race is over by decrementing the waitgroup.
		wg.Done()

		// End this function here so we don't line up another runner.
		return
	}

	// Put the next runner at their starting position.
	go runner(track)

	// Incremember the count of the baton.
	baton++

	// Pass the baton to the next runner by sending it on the channel.
	track <- baton
}
