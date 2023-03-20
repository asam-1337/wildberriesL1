package main

import "math"

type Point struct {
	x float64
	y float64
}

// возвращает указатель на обьект точку
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// рассчитывает расстояние по формуле
func (a *Point) Distance(b *Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func Distance() {

}

func main() {

}
