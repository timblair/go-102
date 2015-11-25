// Create a user type, and an admin type that embeds a user. Create a notifier
// interface, and make your user type satisfy that interface. Write a function
// that accepts a value of the interface type, and ensure it works correctly
// when passed a value of your admin type.
package main

// Add your imports here
import "fmt"

// Define a `notifier` interface
type notifier interface {
	notify()
}

// Create a `user` type, with fields for name and email address.  Ensure your
// type satisfies the Notifier interface.
type user struct {
	name, email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s at %s\n", u.name, u.email)
}

// Create an `admin` type which embeds a user, and has a security level
type admin struct {
	user
	level int
}

// Write a function that accepts a value of your interface and calls the method
// associated with that interface.
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	// Create an admin user
	user := admin{
		user: user{
			name:  "Tim Blair",
			email: "tim@bla.ir",
		},
		level: 10,
	}

	// Send the admin a notification via the function you created
	sendNotification(user)
}
