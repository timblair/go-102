// Create a user type, and an admin type that embeds a user. Create a Notifier
// interface, and make your user type satisfy that interface. Write a function
// that accepts a value of the interface type, and ensure it works correctly
// when passed a value of your admin type.
package main

// Add your imports here
import "fmt"

// Define a Notifier interface
type Notifier interface {
	Notify()
}

// Create a User type, with fields for name and email address.  Ensure your
// type satisfies the Notifier interface.
type User struct {
	Name, Email string
}

func (u User) Notify() {
	fmt.Printf("Sending email to %s at %s\n", u.Name, u.Email)
}

// Create an Admin type which embeds a User, and has a security level
type Admin struct {
	User
	Level int
}

// Write a function that accepts a value of your interface and calls the method
// associated with that interface.
func sendNotification(n Notifier) {
	n.Notify()
}

func main() {
	// Create an admin user
	user := Admin{
		User: User{
			Name:  "Tim Blair",
			Email: "tim@bla.ir",
		},
		Level: 10,
	}

	// Send the admin a notification via the function you created
	sendNotification(user)
}
