package boilerplate

import (
	"log"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	// TODO: Change this
	data []string
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	d.data = data

	return
}

func (d *Day) Part1() interface{} {
	return -1
}

func (d *Day) Part2() interface{} {
	return -1
}

// Things to consider:
// 1. Writing easily testable methods/funcs
// 2. Is parallelism possible?
