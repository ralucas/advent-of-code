package day6

import (
	"log"

	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	data []int
	Days int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, ",")

	d.data = arrayutils.MapToInt(data)

	return
}

func (d *Day) SetDays(days int) {
	d.Days = days
}

func (d *Day) Part1() interface{} {
	state := NewState(d.data)
	for i := 0; i < d.Days; i++ {
		state.Day()
	}

	return state.FishCount()
}

func (d *Day) Part2() interface{} {
	return nil
}
