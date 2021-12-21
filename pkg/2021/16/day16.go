package day16

import (
	"log"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	data string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFile(filepath)

	d.data = data

	return
}

func (d *Day) Part1() interface{} {
	return nil
}

func (d *Day) Part2() interface{} {
	return nil
}
