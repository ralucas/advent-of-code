package day13

import (
	"log"
	"strings"
)

const (
	GridMark         int    = 1
	GridNoMark       int    = 0
	GridStringMark   string = "#"
	GridStringNoMark string = "."
)

type Grid struct {
	points     []Point
	pointMap   map[Point]int
	values     [][]int
	foldValues [][]int
}

func NewGrid(points []Point) *Grid {
	return &Grid{points: points}
}

func (g *Grid) Build() *Grid {
	g.pointMap = make(map[Point]int)

	for _, pt := range g.points {
		g.pointMap[pt] = GridMark
	}

	maxX, maxY := extent(g.pointMap)

	g.values = make([][]int, maxY+1)

	for i := 0; i < maxY+1; i++ {
		g.values[i] = make([]int, maxX+1)
		for j := 0; j < maxX+1; j++ {
			g.values[i][j] = GridNoMark
		}
	}

	for _, pt := range g.points {
		g.values[pt.y][pt.x] = GridMark
	}

	g.foldValues = g.values

	return g
}

func (g *Grid) PointsToGrid() [][]int {
	maxX, maxY := extent(g.pointMap)

	values := make([][]int, maxY+1)

	for i := 0; i < maxY+1; i++ {
		values[i] = make([]int, maxX+1)
		for j := 0; j < maxX+1; j++ {
			values[i][j] = GridNoMark
		}
	}

	for pt, _ := range g.pointMap {
		values[pt.y][pt.x] = GridMark
	}

	return values
}

func (g *Grid) Fold(direction Direction) [][]int {
	switch direction.axis {
	case AxisX:
		return g.foldVertical(direction.value)
	case AxisY:
		return g.foldHorizontal(direction.value)
	default:
		log.Fatal("bad axis")
	}

	return nil
}

func (g *Grid) FoldPoints(direction Direction) {
	switch direction.axis {
	case AxisX:
		g.foldX(direction.value)
	case AxisY:
		g.foldY(direction.value)
	default:
		log.Fatal("bad axis")
	}
}

func (g *Grid) foldX(val int) {
	for point := range g.pointMap {
		if point.x > val {
			rem := point.x % val
			x := val - rem
			if rem == 0 {
				x = 0
			}
			g.pointMap[Point{x, point.y}] = 1
			delete(g.pointMap, point)
		}
	}
}

func (g *Grid) foldY(val int) {
	for point := range g.pointMap {
		if point.y > val {
			rem := point.y % val
			y := val - rem // 10 - 0. 9 - 1, 8 - 2, 7 - 3, 6 - 4
			if rem == 0 {
				y = 0
			}
			g.pointMap[Point{point.x, y}] = 1
			delete(g.pointMap, point)
		}
	}
}

func (g *Grid) foldHorizontal(val int) [][]int {
	foldValues := g.foldValues

	x := 0
	for i := len(g.foldValues) - 1; i > val; i-- {
		for j := 0; j < len(g.foldValues[i]); j++ {
			foldValues[x][j] = g.foldValues[i][j] | foldValues[x][j]
		}
		x += 1
	}

	g.foldValues = foldValues[:val]

	return g.foldValues
}

func (g *Grid) foldVertical(val int) [][]int {
	tmparr := make([][]int, len(g.foldValues))
	foldValues := g.foldValues

	for i := 0; i < len(g.foldValues); i++ {
		y := 0
		for j := len(g.foldValues[i]) - 1; j > val; j-- {
			foldValues[i][y] = g.foldValues[i][j] | foldValues[i][y]
			y += 1
		}
		tmparr[i] = foldValues[i][:val]
	}

	g.foldValues = tmparr

	return g.foldValues
}

func (g *Grid) String(arr [][]int) string {
	var sb strings.Builder
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == GridMark {
				sb.WriteString(GridStringMark)
			} else {
				sb.WriteString(GridStringNoMark)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func extent(points map[Point]int) (int, int) {
	maxX, maxY := 0, 0

	for point, _ := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	return maxX, maxY
}
