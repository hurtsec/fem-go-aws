package main

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func (c *Circle) circumference() float64 {
	return 2 * PI * c.radius
}

func (c *Circle) area() float64 {
	return PI * (c.radius * c.radius)
}

func main() {
	circ := Circle{
		radius: 5,
	}

	fmt.Println(circ.circumference())
	fmt.Println(circ.area())
}
