package main

import "fmt"
import "math"

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	c:= Circle{0, 0, 1}
	fmt.Println(c.x, c.y, c.r)
	//fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	fmt.Println(distance(0,0,1,1))
	r:= Rectangle{0,0,2,2}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))

	m := MultiShape{[]Shape{&c,&r,&r}}
	fmt.Println(m.area())
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y1)
	w := distance(r.x1, r.y1, r.x1, r.y2)
	return l*w
}

func totalArea(shapes ...Shape) (total float64) {
	for _, shape := range shapes {
		total += shape.area()
	}
	return
}

func (m *MultiShape) area() (total float64) {
	for _, shape := range m.shapes {
		total += shape.area()
	}
	return
}