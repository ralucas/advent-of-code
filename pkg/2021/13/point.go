package day13

// x -> col
// y -> row
type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}
