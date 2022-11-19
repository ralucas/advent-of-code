package day9

import (
	"log"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data [][]int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	for _, line := range data {
		lineInts := arrayutil.MapToInt(strings.Split(line, ""))
		d.data = append(d.data, lineInts)
	}

	return
}

func (d *Day) Part1() interface{} {
	grid := NewGrid(d.data)

	totalRiskLevel := 0

	for row := range grid.values {
		for col := range grid.values[row] {
			if grid.isLowPoint(NewPoint(row, col, grid.values[row][col])) {
				totalRiskLevel += grid.values[row][col] + 1
			}
		}
	}

	return totalRiskLevel
}

func (d *Day) Part2() interface{} {
	grid := NewGrid(d.data)

	lowpoints := grid.lowPoints()

	// all surrounding points of a lowpoint
	// are in a basin.
	// search those basin's surrounding points
	// recursively for points that are greater
	// than point.
	// Backtracking?
	basins := make([][]Point, 3)
	for _, point := range lowpoints {
		root := NewNodePoint(point)
		basin := root.BuildBasin(grid)
		lb := len(basin)

		for i, b := range basins {
			if b == nil {
				basins[i] = basin
				break
			}
			if lb > len(b) {
				for j := i + 1; j < len(basins); j++ {
					basins[j] = basins[j-1]
				}
				basins[i] = basin
				break
			}
		}
	}

	ans := 1
	for _, basin := range basins {
		ans *= len(basin)
	}

	return ans
}
