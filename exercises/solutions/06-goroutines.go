// Create two anonymous functions: one that outputs integers from 1 to 100; the
// other from 100 to 1.  Start each function as a goroutine.  Use a WaitGroup
// to ensure that main() doesn't exit until the goroutines are done.
package main

// Add your imports here.
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a new waitgroup.
	var wg sync.WaitGroup

	// Add 2 to the waitgroup; one for each counter function.
	wg.Add(2)

	// Create an anonymous function that counts from 1 to 100, and launch it as
	// a goroutine.
	go func(id string) {
		// Schedule the waitgroup to be decremented when the function exists.
		defer wg.Done()

		// Loop from 1 to 100.
		for i := 1; i <= 100; i++ {
			// Use `time.Sleep` to simulate doing some work so the scheduler
			// has the chance to switch between goroutines.
			time.Sleep(time.Millisecond * 10)

			// Output the current number, prefixed with a value that identifies
			// this function.
			fmt.Println(id, i)
		}
	}("up")

	// Create an anonymous function that counts from 1 to 100, and launch it as
	// a goroutine.
	go func(id string) {
		// Schedule the waitgroup to be decremented when the function exists.
		defer wg.Done()

		// Loop from 100 to 1.
		for i := 100; i >= 1; i-- {
			// Use `time.Sleep` to simulate doing some work so the scheduler
			// has the chance to switch between goroutines.
			time.Sleep(time.Millisecond * 10)

			// Output the current number, prefixed with a value that identifies
			// this function.
			fmt.Println(id, i)
		}
	}("down")

	// Wait until both goroutines have finished before existing main().
	wg.Wait()
}
