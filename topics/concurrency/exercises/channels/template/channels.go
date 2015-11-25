// Let's simulate a track relay race.  Create a channel representing a track,
// and a function representing a runner.  Pass a baton between runners over the
// channel, and end the race when the fourth runner receives the baton.
//
// Template available at: http://play.golang.org/p/H4F9aLKQVA
package main

// Add your imports here.

// Create a new waitgroup.

func main() {
	// Add something to the waitgroup so main() can wait for the race to
	// finish.

	// Create an unbuffered channel to act as the track.  In terms of the
	// channel type, think about how you're going to count the runners.

	// Put the first runner at his starting position by launching the
	// `runner()` function as a goroutine.

	// Give the runner the baton by sending it on the channel.

	// Wait until the race is over.
}

// Create the function representing the runner.  They'll need to be given the
// track channel so we can pass the baton between runners.
func funcName( /* arguments */ ) {
	// Set the runner ready to receive the baton by receiving on the channel.

	// If this runner is the fourth runner then the race is over.
	{
		// Signify the race is over by decrementing the waitgroup.

		// End this function here so we don't line up another runner.
	}

	// Put the next runner at their starting position.

	// Incremember the count of the baton.

	// Pass the baton to the next runner by sending it on the channel.
}
