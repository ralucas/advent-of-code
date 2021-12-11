package day9

type Grid struct {
	values [][]int
	rows   int
	cols   int
}

func NewGrid(data [][]int) *Grid {
	return &Grid{
		values: data,
		rows:   len(data),
		cols:   len(data[0]),
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

func (g *Grid) surroundingPoints(p Point) []Point {
	surrounding := make([]Point, 0)

	offsets := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, offset := range offsets {
		r, c := p.row+offset[0], p.col+offset[1]

		if g.isValidIndex(r, c) {
			surroundingPoint := NewPoint(r, c, g.values[r][c])
			surrounding = append(surrounding, surroundingPoint)
		}
	}

	return surrounding
}

func (g *Grid) isLowPoint(p Point) bool {
	sps := g.surroundingPoints(p)

	for _, sp := range sps {
		if p.value >= sp.value {
			return false
		}
	}

	return true
}

func (g *Grid) lowPoints() []Point {
	lowpoints := make([]Point, 0)

	for row := range g.values {
		for col := range g.values[row] {
			pt := NewPoint(row, col, g.values[row][col])
			if g.isLowPoint(pt) {
				lowpoints = append(lowpoints, pt)
			}
		}
	}

	return lowpoints
}
