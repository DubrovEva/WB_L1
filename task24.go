package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными параметрами x, y и конструктором.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (fp *Point) Distance(sp *Point) float64 {
	return math.Pow(math.Pow(fp.x-sp.x, 2)+math.Pow(fp.y-sp.y, 2), 0.5)
}

func main() {
	firstPoint := NewPoint(3, 4)
	secondPoint := NewPoint(2, 3)

	fmt.Print(firstPoint.Distance(secondPoint)) // 1.4142135623730951
}
