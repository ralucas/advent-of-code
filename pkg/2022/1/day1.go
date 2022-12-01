package day1

import (
	"log"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	data [][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n\n")

	for _, s := range data {
		spl := arrayutil.MapToInt(strings.Split(s, "\n"))
		d.data = append(d.data, spl)
	}

	return
}

func (d *Day) Part1() interface{} {
	var sums []int
	for _, d := range d.data {
		sums = append(sums, mathutil.Sum(d))
	}

	return mathutil.Max(sums...)
}

func (d *Day) Part2() interface{} {
	var sums []int
	for _, d := range d.data {
		sums = append(sums, mathutil.Sum(d))
	}

	sorted := sortutil.QSort(sums)

	var output int
	for i := len(sorted)-1; i > len(sorted)-4; i-- {
		output += sorted[i]
	}

	return output
}
