package structs

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() (res float64) {
	res = r.Height * r.Width
	return
}

func (c Circle) Area() (res float64) {
	res = math.Pi * c.Radius * c.Radius
	return
}

func (t Triangle) Area() (res float64) {
	res = t.Base * t.Height * 0.5
	return
}
