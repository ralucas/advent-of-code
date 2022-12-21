package day12

import (
	"fmt"
	"log"
	"math"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	priorityqueue "github.com/ralucas/advent-of-code/pkg/util/priorityqueue"
	set "github.com/ralucas/advent-of-code/pkg/util/set"
	"github.com/ralucas/advent-of-code/pkg/util/stack"
)

type Day struct {
	grid  [][]byte
	start Location
	end   Location
}

const (
	Start byte = 83
	End   byte = 69
)

// comparable struct
// satisfies pq `Item` interface
type Location struct {
	row     int
	col     int
	value   int
	literal string
}

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d)", l.row, l.col)
}

func (l Location) Neighbors(grid [][]byte) []Location {
	neighbors := make([]Location, 0)

	adj := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, a := range adj {
		x := l.row + a[0]
		y := l.col + a[1]

		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			loc := Location{
				row:     x,
				col:     y,
				value:   int(grid[x][y] - 'a'),
				literal: string(grid[x][y]),
			}

			if grid[x][y] == 'E' {
				loc.value = 26
			}
			if grid[x][y] == 'S' {
				loc.value = 0
			}

			neighbors = append(neighbors, loc)
		}
	}

	return neighbors
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")
	d.grid = arrayutil.MapTo2DGen(data, func(s string, _ int) []byte {
		return []byte(s)
	})

	for row := range d.grid {
		for col := range d.grid[row] {
			if d.grid[row][col] == Start {
				d.start = Location{row, col, 0, "S"}
			}

			if d.grid[row][col] == End {
				d.end = Location{row, col, 26, "E"}
			}
		}
	}

	return
}

func (d *Day) Part1() interface{} {
	path, err := d.Dijkstras(d.start, d.end)
	if err != nil {
		log.Fatal(err)
	}

	return path.Size() - 1
}

func (d *Day) Part2() interface{} {
	possibleStartLocations := []Location{d.start}

	for i := range d.grid {
		for j := range d.grid[i] {
			if d.grid[i][j] == 'a' {
				loc := Location{row: i, col: j, value: 0, literal: "a"}
				possibleStartLocations = append(possibleStartLocations, loc)
			}
		}
	}

	min := math.MaxInt

	for _, start := range possibleStartLocations {
		path, err := d.Dijkstras(start, d.end)
		if err != nil {
			continue
		}

		if path.Size()-1 < min {
			min = path.Size() - 1
		}
	}

	return min
}

// Dijkstras for least cost
func (d *Day) Dijkstras(start, end Location) (*stack.Stack[Location], error) {
	visited := set.New(start)
	distance := make(map[Location]int)
	previous := make(map[Location]Location)

	frontier := priorityqueue.New(priorityqueue.MinPriorityQueue)

	frontier.Insert(start, 0)

	for !frontier.Empty() {
		node, cost := frontier.Pop()

		// todo: this is unsafe...have pq return a generic?
		loc := node.(Location)

		if loc == end {
			break
		}

		neighbors := loc.Neighbors(d.grid)

		for _, neighbor := range neighbors {
			if !visited.Has(neighbor) {
				if neighbor.value-1 <= loc.value {
					// visit it
					visited.Add(neighbor)

					newCost := cost + neighbor.value + 1

					neighborDist, exists := distance[neighbor]

					if !exists || newCost < neighborDist {
						distance[neighbor] = newCost
						previous[neighbor] = loc
						frontier.Insert(neighbor, newCost)
					}
				}
			}
		}
	}

	// ensure previous has end
	if _, ok := previous[end]; !ok {
		return nil, fmt.Errorf("failed to find a path to the end")
	}

	path := stack.New(end)

	for path.Peek() != start {
		path.Push(previous[path.Peek()])
	}

	return path, nil
}
