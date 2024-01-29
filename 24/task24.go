package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// NewPoint это функция-фабрика, реализующая конструктор структуры Point.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(5, 3)
	p2 := NewPoint(10, 20)
	fmt.Printf("%.2f \n", Distance(p1, p2))
}
