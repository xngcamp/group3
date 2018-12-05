package main

import "fmt"

type Graph interface{
	Area() float64
}
//三角形
type Triangle struct {
	hight float64
	wight float64
}
//圆形
type Circle struct {
	r  float64
}
//矩形
type Rectangle struct {
	hight float64
	wight float64
}
//三角形
func (t Triangle) Area() float64{
	return (t.hight*t.wight)/2
}
//圆形
func (c Circle) Area() float64{
	return 3.14*c.r*c.r
}
//矩形
func (r Rectangle) Area() float64{
	return r.hight*r.wight
}

func main(){
	var graph Graph
	graph=Triangle{
		hight :5,
		wight:5,
	}
	fmt.Println("三角形面积为:",graph.Area())
	graph=Circle{
		r:4,
	}
	fmt.Println("圆形面积为:",graph.Area())

	graph=Rectangle{
		hight:5,
		wight:5,
	}
	fmt.Println("矩形面积为:",graph.Area())


}