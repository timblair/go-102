# Go 102 Workshop

## Object-Oriented Programming in Go

### Composition

We've now introduced the building blocks for Go's powerful composition
techniques: custom types, methods, interfaces, and struct embedding.

> "Everyone knows composition is more powerful than inheritance, Go just makes
> this non optional." — Dave Cheney: http://bit.ly/dctlg

Composition is more than just type embedding; it's about declaring and
implementing discrete behaviour, creating types that have a single purpose,
and then using these blocks to build higher level behaviours by composing the
discrete behaviours.

Let's put all these techniques together.  First we'll define a type
representing a location, and two interfaces detailing behaviour.

```go
type Point struct {
	X, Y int
}

type Mover interface {
	MoveTo(p Point)
}

type Firer interface {
	Fire()
}
```

Now we'll define a new type for a vehicle, embedding the location.  We'll also
make this type satisfy one of our interfaces.

```go
type Vehicle struct {
	Point
	Passengers int
}

func (v *Vehicle) MoveTo(p Point) {
	v.Point = p
}
```

Let's also have a weapon, again satisfying the appropriate interface.

```go
type Weapon struct {
	Loaded bool
}

func (w *Weapon) Fire() {
	w.Loaded = false
}
```

With a bit more struct embedding, we can compose our `Vehicle` and `Weapon`
types to create a `Tank`.

```go
type Tank struct {
	Vehicle
	Weapon
}
```

Here we'll introduce a familiar concept with a slight twist: it's type
embedding, but with an interface.  It works in exactly the same way as with
a struct.

```go
type MoverFirer interface {
	Mover
	Firer
}

func moveAndFire(mf MoverFirer, p Point) {
	mf.MoveTo(p)
	mf.Fire()
}
```

And finally, let's wrap everything together.  We'll create a new `Tank`, and
then use our `moveAndFire` function to do just that.

```go
func main() {
	t := &Tank{
		Vehicle{Point{5, 6}, 6},
		Weapon{true},
	}

	moveAndFire(t, Point{10, 20})

	fmt.Printf("Location: %v; Passengers: %d; Loaded: %t\n",
		t.Point, t.Passengers, t.Loaded)

	// Location: {10 20}; Passengers: 6; Loaded: false
}
```

A complete, runnable version of the above example can be found on the [Go
Playground](http://play.golang.org/p/IgLAwNX5ut).
