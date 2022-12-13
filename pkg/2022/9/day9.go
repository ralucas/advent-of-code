package day9

import (
	"fmt"
	"log"
	"strconv"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
)

type Day struct {
	moves      []Move
	allVisited [][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := arrayutil.MapTo2D(fileutil.ReadFileToArray(filepath, "\n"), " ")

	d.moves = make([]Move, len(data))

	for i, item := range data {
		dir := Direction(item[0])
		steps, err := strconv.Atoi(item[1])
		if err != nil {
			log.Fatal(err)
		}

		d.moves[i] = Move{direction: dir, steps: steps}
	}

	return
}

func printVisited(visited map[Position]int) {
	xs := make([]int, len(visited))
	ys := make([]int, len(visited))

	i := 0
	for pos := range visited {
		xs[i] = pos.x
		ys[i] = pos.y
		i += 1
	}

	minX, maxX := mathutil.Extent(xs)
	width := maxX - minX + 1
	minY, maxY := mathutil.Extent(ys)
	height := maxY - minY + 1

	graph := make([][]int, height)
	for i := range graph {
		graph[i] = make([]int, width)
	}

	for pos := range visited {
		x, y := pos.x-minX, pos.y-minY
		graph[y][x] = 1
	}

	for i := len(graph) - 1; i >= 0; i-- {
		for j := range graph[i] {
			fmt.Printf("%d ", graph[i][j])
		}
		fmt.Print("\n")
	}
}

func (d *Day) AllVisited() [][]int {
	return d.allVisited
}

func (d *Day) Part1() interface{} {
	H := NewStartPosition()
	T := NewStartPosition()

	visited := make(map[Position]int)
	visited[Position{T.x, T.y}] = 1
	d.allVisited = [][]int{{T.x, T.y}}

	for _, move := range d.moves {
		done := false
		for !done {
			move, done = H.MakeStep(move)
			xDiff, yDiff := H.Difference(T)
			mfd := MovesFromDifference(xDiff, yDiff)
			for len(mfd) > 1 {
				T.MakeStep(mfd[0])
				key := Position{T.x, T.y}
				visited[key] = 1
				d.allVisited = append(d.allVisited, []int{key.x, key.y})
				mfd = mfd[1:]
			}
		}
	}

	return len(visited)
}

func (d *Day) Part2() interface{} {
	H := &Position{}
	knots := make([]*Position, 9)
	for i := range knots {
		knots[i] = &Position{}
	}

	visited := make(map[Position]int)
	visited[Position{0, 0}] = 1
	d.allVisited = [][]int{{0, 0}}

	for _, move := range d.moves {
		done := false
		for !done {
			move, done = H.MakeStep(move)
			leader := H
			for i, knot := range knots {
				xDiff, yDiff := leader.Difference(knot)
				mfd := MovesFromDifference(xDiff, yDiff)
				for len(mfd) > 1 {
					knot.MakeStep(mfd[0])
					mfd = mfd[1:]
					// if it's the tail
					if i == len(knots)-1 {
						// key := fmt.Sprintf("%d%d", knot.x, knot.y)
						key := Position{knot.x, knot.y}
						visited[key] = 1
						d.allVisited = append(d.allVisited, []int{knot.x, knot.y})
					}
				}

				leader = knot
			}
		}
	}

	// printVisited(visited)
	return len(visited)
}
