package day1

import (
	"log"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data []int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	d.data = arrayutil.MapToInt(data)

	return
}

func (d *Day) Part1() interface{} {
	count := 0

	if len(d.data) <= 1 {
		return count
	}

	for i := 1; i < len(d.data); i++ {
		if d.data[i] > d.data[i-1] {
			count += 1
		}
	}

	return count
}

func (d *Day) Part2() interface{} {
	count := 0

	if len(d.data) <= 1 {
		return count
	}

	for i := 3; i < len(d.data); i++ {
		if d.data[i] > d.data[i-3] {
			count += 1
		}
	}

	return count
}
