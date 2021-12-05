package day5

import (
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
	points     []Point
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func NewLine(start, end Point) Line {
	return Line{
		start:  start,
		end:    end,
		points: buildPoints(start, end),
	}
}

func (l Line) Points() []Point {
	return l.points
}

func buildPoints(start, end Point) []Point {
	var points []Point

	if start.x == end.x {
		points = append(points, start)
		diff := mathutils.Abs(end.y - start.y)
		y := start.y
		for i := 0; i < diff-1; i++ {
			if end.y > start.y {
				y += 1
				points = append(points, NewPoint(start.x, y))
			} else {
				y -= 1
				points = append(points, NewPoint(start.x, y))
			}
		}
		points = append(points, end)
	}

	if start.y == end.y {
		points = append(points, start)
		diff := mathutils.Abs(end.x - start.x)
		x := start.x
		for i := 0; i < diff-1; i++ {
			if end.x > start.x {
				x += 1
				points = append(points, NewPoint(x, start.y))
			} else {
				x -= 1
				points = append(points, NewPoint(x, start.y))
			}
		}
		points = append(points, end)
	}

	return points
}
