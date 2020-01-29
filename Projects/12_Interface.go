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

//Interface Definition
type Shape interface {
	Area() float64
	Perimeter() float64
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

//Common function to print Area and perimeter
func ShapeDetails(s []Shape) {
	area := 0.00
	perimeter := 0.00
	for _, t := range s {
		area = area + t.Area()
		perimeter = perimeter + t.Perimeter()
		fmt.Printf("\nArea of %T : %f", t, t.Area())  //Print the area of the shape
		fmt.Printf("\nPerimeter of %T : %f", t, t.Perimeter())	//Print the perimeter of the shape
	}
	fmt.Println("\n--------Total--------")
	fmt.Printf("\nTotal Area of all Shapes: %f", area)
	fmt.Printf("\nTotal Perimeter of all Shapes: %f", perimeter)
}

func main() {
	//Decalre different shapes with thier measurements
	s1 := Rectangle{5, 8}
	s2 := Circle{4}
	s3 := Rectangle{50, 80}
	s4 := Circle{5}
	s5 := Rectangle{1.2, 1.5}
	s6 := Circle{6}

	//Declare Interface
	I1 := []Shape{s1, s2, s3, s4, s5, s6}

	//Invoke the Shape details function using Interface
	ShapeDetails(I1)
}
