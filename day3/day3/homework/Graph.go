package main

import "fmt"

type Graph interface {
	Area() float64
}

//三角形
type Triangle struct {
	width float64
	high  float64
}

func (t Triangle) Area() float64 {
	return (t.high * t.width) / 2
}

//圆
type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.r * c.r
}

//矩形
type Rectangle struct {
	long  float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.long
}

func main() {
	var g Graph
	g = Triangle{
		10,
		5,
	}
	fmt.Printf("三角形的面积为%v\n", getArea(g))

	g = Circle{
		5,
	}
	fmt.Printf("圆的面积为%v\n", getArea(g))

	g = Rectangle{
		10,
		5,
	}
	fmt.Printf("矩形的面积为%v\n", getArea(g))
}
func getArea(g Graph) float64 {
	return g.Area()
}
