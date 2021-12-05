package day5

import (
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Point struct {
	x, y int
}

type LineBuilder struct {
	*Line
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

func NewLineBuilder(start, end Point) *LineBuilder {
	return &LineBuilder{
		&Line{
			start: start,
			end:   end,
		},
	}
}

func NewLine(start, end Point) *Line {
	return &Line{
		start: start,
		end:   end,
	}
}

func (l *LineBuilder) BuildWith() *LineBuilder {
	if len(l.points) > 0 {
		l.points = nil
	}

	return l
}

func (l *LineBuilder) BuildLine() *Line {
	return l.Line
}

func (l *LineBuilder) Horizontal() *LineBuilder {
	if l.start.x == l.end.x {
		l.points = append(l.points, l.start)
		diff := mathutils.Abs(l.end.y - l.start.y)
		y := l.start.y
		for i := 0; i < diff-1; i++ {
			if l.end.y > l.start.y {
				y += 1
				l.points = append(l.points, NewPoint(l.start.x, y))
			} else {
				y -= 1
				l.points = append(l.points, NewPoint(l.start.x, y))
			}
		}
		l.points = append(l.points, l.end)
	}

	return l
}

func (l *LineBuilder) Vertical() *LineBuilder {
	if l.start.y == l.end.y {
		l.points = append(l.points, l.start)
		diff := mathutils.Abs(l.end.x - l.start.x)
		x := l.start.x
		for i := 0; i < diff-1; i++ {
			if l.end.x > l.start.x {
				x += 1
				l.points = append(l.points, NewPoint(x, l.start.y))
			} else {
				x -= 1
				l.points = append(l.points, NewPoint(x, l.start.y))
			}
		}
		l.points = append(l.points, l.end)
	}

	return l
}

func (l *LineBuilder) Diagonal() *LineBuilder {
	if l.start.x != l.end.x && l.start.y != l.end.y {
		l.points = append(l.points, l.start)

		diffX := mathutils.Abs(l.end.x - l.start.x)
		diffY := mathutils.Abs(l.end.y - l.start.y)

		// must be diagonal possible
		if diffX == diffY {
			x := l.start.x
			y := l.start.y
			for i := 0; i < diffX-1; i++ {
				if l.end.x > l.start.x {
					x += 1
				} else {
					x -= 1
				}
				if l.end.y > l.start.y {
					y += 1
				} else {
					y -= 1
				}
				l.points = append(l.points, NewPoint(x, y))
			}
			l.points = append(l.points, l.end)
		}
	}

	return l
}

func (l *Line) Points() []Point {
	return l.points
}
