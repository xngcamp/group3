package main

import "fmt"

type Graph interface {
	Area() float64
}

type Triangle struct {
	Button float64
	Heighth float64
}

type Circle struct {
	R float64
}

type Rectangle struct {
	Longth float64
	Weight float64
}

func (triangle *Triangle) Area() float64 {
	s := (triangle.Button * triangle.Heighth)/2
	return s
}

func (circle *Circle) Area() float64 {
	s := circle.R * circle.R * 3.14
	return s
}

func (rectangle *Rectangle) Area() float64 {
	s := rectangle.Longth * rectangle.Weight
	return s
}

func main() {
	var triangle1 Triangle
	triangle1.Heighth = 10
	triangle1.Button = 15
	fmt.Println(triangle1.Area())

	var circle1 Circle
	circle1.R = 17.0
	fmt.Println(circle1.Area())

	var rectangle1 Rectangle
	rectangle1.Weight = 2.5
	rectangle1.Longth = 7.0
	fmt.Println(rectangle1.Area())
}


