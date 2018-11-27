# Problems

- What's the difference between a method and a function?

A method is a function that's called on a type using dot notation.

- Why would you use an embedded anonymous field instead of a normal named field?

Embedded anonymous fields denote `is-a` relationships where named fields denote `has-a` relationships.

- Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.

```golang
type Shape interface {
  area() float64
  perimeter() float64
}

func (c *Circle) perimeter() float64 {
	return math.Pi * (c.r * 2)
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return r.x1 + r.x2 + r.x3 + r.x4
}
```