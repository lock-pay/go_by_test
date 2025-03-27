package structs

import "math"

func Perimeter(shape Rectangle) float64 {
	return 2 * (shape.Height + shape.Width)
}

type Shape interface {
	Area() float64
}

func Area(shape Rectangle) float64 {
	return shape.Width * shape.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius int
}

type Triangle struct {
	Base   int
	Height int
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2.0)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (float64(t.Base) * float64(t.Height) * 0.5)
}
