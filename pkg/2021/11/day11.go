package day11

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

	vals := make([][]int, len(data))
	for i, line := range data {
		vals[i] = arrayutils.MapToInt(strings.Split(line, ""))
	}

	d.data = vals

	return
}

func (d *Day) Part1() interface{} {
	grid := NewGrid(d.data)

	for i := 0; i < 100; i++ {
		grid.Step()
	}

	return grid.FlashCount()
}

func (d *Day) Part2() interface{} {
	grid := NewGrid(d.data)

	gridSize := grid.cols * grid.rows

	step := 1
	for {
		grid.Step()

		if gridSize == grid.StepFlashCount() {
			return step
		}

		step++
	}
}
