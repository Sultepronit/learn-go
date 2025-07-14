package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.
// to implement interface is to implement all it's methods

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perim() float64 {
	return s.side * 4
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return c.radius * 2 * math.Pi
}

func measure(f geometry) {
	if c, ok := f.(circle); ok {
		fmt.Println("circle with radius:", c.radius)
	} else if s, ok := f.(square); ok {
		fmt.Println("square with side:", s.side)
	}

	fmt.Println(f.area())
	fmt.Println(f.perim())
}

func main() {
	s := square{5}
	c := circle{10}

	measure(s)
	measure(c)
}
