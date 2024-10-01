package structs

import "math"

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimiter() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimiter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimiter() float64 {
	return 2 * r.Height * r.Width
}
