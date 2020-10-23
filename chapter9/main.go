package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type MultiShape struct { //Interfaces can also be used as fields
	shapes []Shape
}

type Shape interface {
	area() float64
}

func (r *Rectangle) area() float64 {
	l := 15.
	w := 12.
	return l * w
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//
func circleAreaWithPointer(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

//(c *Circle) this is called method and allow us to call area from a type Circle like c.area()
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	//c := Circle{x: 0, y: 0, r: 5} //other way to init the values
	c := Circle{0, 0, 5}

	fmt.Println(circleArea(c))

	fmt.Println(circleAreaWithPointer(&c))

	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))
}
