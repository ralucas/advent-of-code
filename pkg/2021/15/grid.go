package day15

import (
	"fmt"
	"strings"
)

type GridBuilder struct {
	*Grid
}

func NewGridBuilder(data [][]int) *GridBuilder {
	return &GridBuilder{
		Grid: &Grid{
			values: data,
			rows:   len(data),
			cols:   len(data[0]),
		},
	}
}

type Grid struct {
	values [][]int
	nodes  [][]*Node
	rows   int
	cols   int
}

func NewGrid(data [][]int) *Grid {
	g := &Grid{
		values: data,
	}

	return g
}

func (gb *GridBuilder) Expand(multiplier int) *GridBuilder {
	expanded := make([][]int, multiplier*len(gb.values))

	for i := range expanded {
		expanded[i] = make([]int, multiplier*len(gb.values[0]))
	}

	w, _ := len(gb.values[0]), len(gb.values)

	for i := range gb.values {
		for j := range gb.values[i] {
			expanded[i][j] = gb.values[i][j]
			for x := 1; x < multiplier; x++ {
				r, c, cprev := i, j+(w*x), j+(w*(x-1))
				expanded[r][c] = expanded[r][cprev] + 1
				if expanded[r][c] > 9 {
					expanded[r][c] = 1
				}
			}
		}
	}

	for i := 0; i < multiplier-1; i++ {
		ph := len(gb.values) * i
		h, hend := len(gb.values)*(i+1), len(gb.values)*(i+2)

		i := ph
		for h < hend {
			for j := 0; j < len(expanded[i]); j++ {
				expanded[h][j] = expanded[i][j] + 1
				if expanded[h][j] > 9 {
					expanded[h][j] = 1
				}
			}
			i++
			h++
		}
	}

	gb.Grid.values = expanded
	gb.Grid.rows = len(expanded)
	gb.Grid.cols = len(expanded[0])

	return gb
}

func (gb *GridBuilder) Build() *Grid {
	gb.Grid.nodes = make([][]*Node, len(gb.values))

	for row := range gb.values {
		gb.Grid.nodes[row] = make([]*Node, len(gb.values[row]))
		for col := range gb.values[row] {
			gb.Grid.nodes[row][col] = NewNode(row, col)
		}
	}

	for row := range gb.Grid.nodes {
		for _, node := range gb.Grid.nodes[row] {
			snodes := gb.surroundingNodes(node)
			for _, sn := range snodes {
				node.SetNeighbor(sn, gb.values[sn.row][sn.col])
			}
		}
	}

	return gb.Grid
}

func (g *Grid) Values() [][]int {
	return g.values
}

func (g *Grid) NodeValue(row, col int) int {
	return g.values[row][col]
}

func (gb *GridBuilder) isValidIndex(row, col int) bool {
	return row >= 0 &&
		row < gb.rows &&
		col >= 0 &&
		col < gb.cols
}

func (gb *GridBuilder) surroundingNodes(p *Node) []*Node {
	surrounding := make([]*Node, 0)

	offsets := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, offset := range offsets {
		r, c := p.row+offset[0], p.col+offset[1]

		if gb.isValidIndex(r, c) {
			surroundingNode := gb.Grid.nodes[r][c]
			surrounding = append(surrounding, surroundingNode)
		}
	}

	return surrounding
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i := 0; i < len(g.values); i++ {
		for j := 0; j < len(g.values[i]); j++ {
			s := fmt.Sprintf("%d", g.values[i][j])
			sb.WriteString(s)
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
