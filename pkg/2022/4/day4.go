package day4

import (
	"log"
	"strconv"
	"strings"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	// bitutil "github.com/ralucas/advent-of-code/pkg/util/bit"
	// mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	// sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	data [][][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}

	data := fileutil.ReadFileToArray(filepath, "\n")
	data2 := arrayutil.MapTo2D(data, ",")
	data3 := make([][][]int, len(data2))

	for i := range data2 {
		data3[i] = make([][]int, 2)
		for j, s := range data2[i] {
			ss := strings.Split(s, "-")

			vi := make([]int, len(ss))

			var err error

			for x := range ss {
				vi[x], err = strconv.Atoi(ss[x])
				if err != nil {
					log.Fatalf("failed to convert %v", err)
				}
			}

			data3[i][j] = vi
		}
	}

	d.data = data3

	return
}

func (d *Day) Part1() interface{} {
	count := 0

	for i := range d.data {
		first, second := d.data[i][0], d.data[i][1]
		a, b := first[0], first[1]
		x, y := second[0], second[1]

		if a <= x && b >= y {
			count += 1
		} else if x <= a && y >= b {
			count += 1
		}
	}

	return count
}

func (d *Day) Part2() interface{} {
	count := 0

	for i := range d.data {
		first, second := d.data[i][0], d.data[i][1]
		a, b := first[0], first[1]
		x, y := second[0], second[1]

		if a >= x && a <= y {
			count += 1
		} else if b <= y && b >= x {
			count += 1
		} else if x >= a && x <= b {
			count += 1
		} else if y <= b && y >= a {
			count += 1
		}
	}

	return count
}
