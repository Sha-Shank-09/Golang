package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	length, breadth float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{radius: 7.5}
	rectangle := Rectangle{length: 5, breadth: 2}

	fmt.Printf("Area of the circle is %f\n", getArea(circle))
	fmt.Printf("Area of Rectangle is %f\n", getArea(rectangle))
}
