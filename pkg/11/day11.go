package day11

import (
	"log"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	data [][]string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	out := utils.MapTo2D(data, ",")

	d.data = out

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
