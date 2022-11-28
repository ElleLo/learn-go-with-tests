package structs

import "math"

// Rectangle and Circle has a method called Area which returns a float64,
// this satisfies the Shape interface
// String does not have such a method, so it doesn't satisfy the interface
type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Method syntax: func (receiverName ReceiverType) Methodname(args) ReturnType
// Go convention: receivedName = 1st letter of ReceiverType
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
