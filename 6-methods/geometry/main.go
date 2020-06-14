package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// traditional function
func distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{X: 3, Y: 4}
	q := Point{1, 2}
	fmt.Println(distance(p, q))
	fmt.Println(p.distance(q))
}
