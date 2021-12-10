package day9

import (
	"log"
	"strings"

	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	data [][]int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, "\n")

	for _, line := range data {
		lineInts := arrayutils.MapToInt(strings.Split(line, ""))
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

	return nil
}
