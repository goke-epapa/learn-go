package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	Base   float64
	Height float64
}

type square struct {
	Side float64
}

func main() {
	tri := triangle{
		Base:   4,
		Height: 10,
	}
	printArea(tri)

	sq := square{Side: 4}
	printArea(sq)
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.Base * tr.Height
}

func (sq square) getArea() float64 {
	return math.Pow(sq.Side, 2)
}

func printArea(s shape) {
	fmt.Printf("Area %f", s.getArea())
}
