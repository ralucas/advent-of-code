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

	// if maxX%2 == 0 {
	// 	maxX += 1
	// }

	// if maxY%2 == 0 {
	// 	maxY += 1
	// }

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
	foldValues := g.values

	x := 0
	for i := len(g.values) - 1; i > val; i-- {
		for j := 0; j < len(g.values[i]); j++ {
			foldValues[x][j] = g.values[i][j] | foldValues[x][j]
		}
		x += 1
	}

	fv := foldValues[:val]

	return fv
}

func (g *Grid) foldVertical(val int) [][]int {
	output := make([][]int, len(g.values))
	foldValues := g.values

	for i := 0; i < len(g.values); i++ {
		y := 0
		for j := len(g.values[i]) - 1; j > val; j-- {
			foldValues[i][y] = g.values[i][j] | foldValues[i][y]
			y += 1
		}
		output[i] = foldValues[i][:val]
	}

	return output
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
