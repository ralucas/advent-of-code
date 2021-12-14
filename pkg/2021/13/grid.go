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
	points     []*Point
	values     [][]int
	foldValues [][]int
}

func NewGrid(points []*Point) *Grid {
	return &Grid{points: points}
}

func (g *Grid) Build() *Grid {
	maxX, maxY := extent(g.points)

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

func extent(points []*Point) (int, int) {
	maxX, maxY := 0, 0

	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	return maxX, maxY
}
