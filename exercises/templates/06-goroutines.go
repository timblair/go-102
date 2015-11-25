// Create two anonymous functions: one that outputs integers from 1 to 100; the
// other from 100 to 1.  Start each function as a goroutine.  Use a WaitGroup
// to ensure that main() doesn't exit until the goroutines are done.
//
// Template available at: http://play.golang.org/p/...
package main

// Add your imports here.

func main() {
	// Create a new waitgroup.

	// Add 2 to the waitgroup; one for each counter function.

	// Create an anonymous function that counts from 1 to 100, and launch it as
	// a goroutine.
	{
		// Schedule the waitgroup to be decremented when the function exists.

		// Loop from 1 to 100.
		{
			// Use `time.Sleep` to simulate doing some work so the scheduler
			// has the chance to switch between goroutines.

			// Output the current number, prefixed with a value that identifies
			// this function.
		}
	}

	// Create an anonymous function that counts from 1 to 100, and launch it as
	// a goroutine.
	{
		// Schedule the waitgroup to be decremented when the function exists.

		// Loop from 100 to 1.
		{
			// Use `time.Sleep` to simulate doing some work so the scheduler
			// has the chance to switch between goroutines.

			// Output the current number, prefixed with a value that identifies
			// this function.
		}
	}

	// Wait until both goroutines have finished before existing main().
}
