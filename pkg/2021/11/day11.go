package day11

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

	vals := make([][]int, len(data))
	for i, line := range data {
		vals[i] = arrayutil.MapToInt(strings.Split(line, ""))
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
