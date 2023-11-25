package day13

import (
	"log"
	"strings"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	// arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	// bitutil "github.com/ralucas/advent-of-code/pkg/util/bit"
	// mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	// sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	pairs [][][]*Element
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n\n")

	for _, lines := range data {
		slines := strings.Split(lines, "\n")
		pair := make([][]*Element, 2)
		for i, line := range slines {
			list, err := Parse(line)
			if err != nil {
				log.Fatal(err)
			}

			pair[i] = list
		}
		d.pairs = append(d.pairs, pair)
	}

	return
}

func (d *Day) Part1() interface{} {
	for _, pair := range d.pairs {
		left, right := pair[0], pair[1]

		for i := range left {
			lel := left[i]
			if lel.IsLeaf() {
				
			}
		}
	}
}

func (d *Day) Part2() interface{} {
	return nil
}
