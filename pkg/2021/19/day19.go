package day19

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
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
	data := fileutil.ReadFileToArray(filepath, "\n")

	d.data = data

	return
}

func (d *Day) Part1() interface{} {
	return nil
}

func (d *Day) Part2() interface{} {
	return nil
}
