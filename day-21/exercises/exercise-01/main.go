package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	var shape Shape

	shape = Rectangle{Length: 10, Width: 5}
	fmt.Printf("Rectangle Area: %.2f\n", shape.Area())

	shape = Circle{Radius: 7}
	fmt.Printf("Circle Area: %.2f\n", shape.Area())
}
