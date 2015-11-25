// Declare a new struct type to hold information about a tennis player,
// including the number of matches played and the number won.  Add a method to
// this type that calculates the win ratio for the player.  Create a new
// player, and output the win ratio for them.
package main

// Add your imports here.
import "fmt"

// Declare a struct type `player` to maintain information about a player.
type player struct {
	name    string
	matches int
	wins    int
}

// Declare a method that calculates the win ratio for the player.  Note that
// you'll likely need to convert one or more values to floats, which can be
// done like: float32(intValue).
func (p player) winRatio() float32 {
	return float32(p.wins) / float32(p.matches)
}

func main() {
	// Create a new player, and output their win ratio.
	p1 := player{"Serena", 20, 19}
	fmt.Println(p1.name, p1.winRatio())

	// If you're feeling adventurous, try creating a slice of multiple players
	// and iterating over the slice, displaying the player name and win ratio.
	p2 := []player{
		{"Venus", 17, 8},
		{"Martina", 44, 30},
		{"Bjorn", 23, 13},
		{"Boris", 34, 24},
		{"Andy", 32, 4},
	}

	for _, p := range p2 {
		fmt.Println(p.name, p.winRatio())
	}
}
