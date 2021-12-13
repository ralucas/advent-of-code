package day11

import "fmt"

type Grid struct {
	values [][]int
	points [][]*Point
	rows   int
	cols   int
}

func NewGrid(data [][]int) *Grid {
	g := &Grid{
		values: data,
		rows:   len(data),
		cols:   len(data[0]),
	}

	g.buildPoints()

	return g
}

func (g *Grid) buildPoints() {
	g.points = make([][]*Point, len(g.values))

	for row := range g.values {
		g.points[row] = make([]*Point, len(g.values[row]))
		for col, val := range g.values[row] {
			g.points[row][col] = NewPoint(row, col, val)
		}
	}

	for row := range g.points {
		for _, point := range g.points[row] {
			sps := g.surroundingPoints(point)
			for _, sp := range sps {
				point.RegisterObserver(sp)
			}
		}
	}
}

func (g *Grid) Values() [][]int {
	return g.values
}

func (g *Grid) isValidIndex(row, col int) bool {
	return row >= 0 &&
		row < g.rows &&
		col >= 0 &&
		col < g.cols
}

func (g *Grid) surroundingPoints(p *Point) []*Point {
	surrounding := make([]*Point, 0)

	offsets := [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{-1, -1},
		{0, -1},
	}

	for _, offset := range offsets {
		r, c := p.row+offset[0], p.col+offset[1]

		if g.isValidIndex(r, c) {
			surroundingPoint := g.points[r][c]
			surrounding = append(surrounding, surroundingPoint)
		}
	}

	return surrounding
}

func (g *Grid) Step() {
	g.resetPoints()

	for _, points := range g.points {
		for _, point := range points {
			point.Add(1)
		}
	}
}

func (g *Grid) FlashCount() int {
	total := 0

	for _, points := range g.points {
		for _, point := range points {
			total += point.FlashCount()
		}
	}

	return total
}

func (g *Grid) resetPoints() {
	for _, points := range g.points {
		for _, point := range points {
			point.Reset()
		}
	}
}

func (g *Grid) StepFlashCount() int {
	count := 0
	for row := range g.points {
		for _, point := range g.points[row] {
			if point.DidFlash() {
				count += 1
			}
		}
	}

	return count
}

func (g *Grid) Print() {
	for _, points := range g.points {
		for _, point := range points {
			fmt.Print(point.value.current)
		}
		fmt.Print("\n")
	}
}
