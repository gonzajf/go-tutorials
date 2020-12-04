package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {

	s := square{
		sideLength: 4,
	}

	t := triangle{
		base:   2,
		height: 4,
	}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
