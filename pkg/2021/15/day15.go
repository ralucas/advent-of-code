package day15

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/fatih/color"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	builder *GridBuilder
	grid    *Grid
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	gridData := make([][]int, 0)
	for _, line := range data {
		gridData = append(gridData, arrayutil.MapToInt(strings.Split(line, "")))
	}

	d.builder = NewGridBuilder(gridData)

	return
}

func (d *Day) Builder() *GridBuilder {
	return d.builder
}

func (d *Day) Grid() *Grid {
	return d.grid
}

func (d *Day) Part1() interface{} {
	d.grid = d.builder.Build()
	start := d.grid.nodes[0][0]
	end := d.grid.nodes[len(d.grid.nodes)-1][len(d.grid.nodes[0])-1]

	dist, _ := d.Dijkstras(start, end)

	// d.PrintPath(start, end, prev)

	return dist[end]
}

func (d *Day) Part2() interface{} {
	d.grid = d.builder.Expand(5).Build()
	start := d.grid.nodes[0][0]
	end := d.grid.nodes[len(d.grid.nodes)-1][len(d.grid.nodes[0])-1]

	dist, _ := d.Dijkstras(start, end)

	// d.PrintPath(start, end, prev)

	return dist[end]
}

func (d *Day) PrintPath(start, end *Node, prev map[*Node]*Node) {
	key := end
	path := make(map[*Node]int)
	path[end] = 1

	for {
		p := prev[key]
		if p != nil {
			path[p] = 1
		}
		if p == start {
			path[p] = 1
			break
		}
		key = p
	}

	c := color.New(color.FgHiRed).Add(color.Underline)

	for i := range d.grid.nodes {
		for j := range d.grid.nodes[i] {
			n := d.grid.nodes[i][j]
			if _, ok := path[n]; ok {
				s := fmt.Sprintf("%d", d.grid.values[i][j])
				c.Print(s)
			} else {
				fmt.Printf("%d", d.grid.values[i][j])
			}
		}
		fmt.Print("\n")
	}

}

func (d *Day) Dijkstras(start, end *Node) (map[*Node]int, map[*Node]*Node) {
	distance := make(map[*Node]int)
	previous := make(map[*Node]*Node)

	for i := range d.grid.nodes {
		for j := range d.grid.nodes[i] {
			n := d.grid.nodes[i][j]
			distance[n] = math.MaxInt
			previous[n] = nil
		}
	}

	distance[start] = 0

	pq := NewMinPriorityQueue()

	pq.Insert(start, 0)

	visited := make(map[*Node]int)

	visited[start] = 1

	for !pq.Empty() {
		node, cost := pq.Pop()

		if node == end {
			return distance, previous
		}

		for neighbor, neighborCost := range node.Neighbors() {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = 1

				newCost := cost + neighborCost

				if newCost < distance[neighbor] {
					distance[neighbor] = newCost
					previous[neighbor] = node
					pq.Insert(neighbor, newCost)
				}
			}
		}
	}

	return distance, previous
}
