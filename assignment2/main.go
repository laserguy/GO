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
	T := triangle{
		height: 10,
		base:   1.9,
	}

	S := square{
		sideLength: 2.9,
	}

	printArea(T)
	printArea(S)
}

func printArea(sh shape) {
	area := sh.getArea()
	fmt.Println(area)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}