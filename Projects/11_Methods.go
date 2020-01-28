package main

import (
	"fmt"
	"math"
)

//Structure Definition
type Rectangle struct {
	L, B float64
}

type Circle struct {
	R float64
}

//Method Definition
func (r Rectangle) Area() float64 {
	return (r.L * r.B)
}

func (c Circle) Area() float64 {
	return (math.Pi * c.R * c.R)
}

func (r Rectangle) Perimeter() float64 {
	return (2 * r.L) + (2 * r.B)
}

func (c Circle) Perimeter() float64 {
	return (math.Pi * 2 * c.R)
}

func main() {
	s1 := Rectangle{5, 8}
	var s2 Circle
	s2.R = 4

	fmt.Println("--------Rectangle--------")
	fmt.Println("Area of Rectangle with Length", s1.L, "and Breadth", s1.B, ": ", s1.Area())
	fmt.Println("Perimeter of Rectangle with Length", s1.L, "and Breadth", s1.B, ": ", (&s1).Perimeter())

	fmt.Println("\n-------- Circle --------")
	fmt.Println("Area of Circle with Radius", s2.R, ":", s2.Area())
	fmt.Println("Perimeter of Circle with Radius", s2.R, ":", s2.Perimeter())
}
