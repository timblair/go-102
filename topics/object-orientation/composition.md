# Object-Oriented Programming

## Composition

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
type point struct {
	x, y int
}

type mover interface {
	moveTo(p point)
}

type firer interface {
	fire()
}
```

Now we'll define a new type for a vehicle, embedding the location.  We'll also
make this type satisfy one of our interfaces.

```go
type vehicle struct {
	point
	passengers int
}

func (v *vehicle) moveTo(p point) {
	v.point = p
}
```

Let's also have a weapon, again satisfying the appropriate interface.

```go
type weapon struct {
	loaded bool
}

func (w *weapon) fire() {
	w.loaded = false
}
```

With a bit more struct embedding, we can compose our `vehicle` and `weapon`
types to create a `tank`.

```go
type tank struct {
	vehicle
	weapon
}
```

Here we'll introduce a familiar concept with a slight twist: it's type
embedding, but with an interface.  It works in exactly the same way as with
a struct.

```go
type moverFirer interface {
	mover
	firer
}

func moveAndFire(mf moverFirer, p point) {
	mf.moveTo(p)
	mf.fire()
}
```

And finally, let's wrap everything together.  We'll create a new `tank`, and
then use our `moveAndFire` function to do just that.

```go
func main() {
	t := &tank{
		vehicle{point{5, 6}, 6},
		weapon{true},
	}

	moveAndFire(t, point{10, 20})

	fmt.Printf("Location: %v; Passengers: %d; Loaded: %t\n",
		t.point, t.passengers, t.loaded)

	// Location: {10 20}; Passengers: 6; Loaded: false
}
```

A complete, runnable version of the above example can be found on the [Go
Playground](http://play.golang.org/p/k145c72ZV4).
