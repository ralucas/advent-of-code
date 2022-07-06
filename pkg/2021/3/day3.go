package day3

import (
	"log"
	"strings"
	"sync"

	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	bitutils "github.com/ralucas/advent-of-code/pkg/utils/bit"
	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	// TODO: Change this
	data [][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}

	data := fileutils.ReadFileToArray(filepath, "\n")

	for _, s := range data {
		vs := strings.Split(s, "")
		vi := arrayutils.MapToInt(vs)
		d.data = append(d.data, vi)
	}

	return
}

func (d *Day) counts(vvi [][]int) [][]int {
	counts := make([][]int, len(vvi[0]))
	for i := range counts {
		counts[i] = []int{0, 0}
	}

	var wg sync.WaitGroup

	for i := range vvi {
		wg.Add(1)
		go func(idx int) {
			for j, v := range vvi[idx] {
				if v == 0 {
					counts[j][0] += 1
				} else {
					counts[j][1] += 1
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	return counts
}

func (d *Day) Part1() interface{} {
	counts := d.counts(d.data)

	gamma := make([]int8, len(counts))
	epsilon := make([]int8, len(counts))

	for i, v := range counts {
		if v[0] > v[1] {
			gamma[i] = int8(0)
			epsilon[i] = int8(1)
		} else {
			gamma[i] = int8(1)
			epsilon[i] = int8(0)
		}
	}

	return bitutils.Btoi(gamma) * bitutils.Btoi(epsilon)
}

func maxIndex(vi []int) int {
	if len(vi) == 0 {
		return -1
	}

	mi := 0
	max := vi[0]

	for i, v := range vi {
		if v >= max {
			max = v
			mi = i
		}
	}

	return mi
}

func (d *Day) filterRatings(counts [][]int, max bool) []int {
	ratings := make([][]int, len(d.data))
	for i := range d.data {
		ratings[i] = make([]int, len(d.data[i]))
		copy(ratings[i], d.data[i])
	}

	for i := 0; i < len(counts); i++ {
		m := maxIndex(counts[i])
		newRatings := arrayutils.FilterInt2D(
			ratings,
			func(vi []int) bool {
				if max {
					return vi[i] == m
				} else {
					return vi[i] != m
				}
			},
		)

		if len(newRatings) != 0 {
			ratings = newRatings
		}

		if len(ratings) == 1 {
			break
		}

		counts = d.counts(ratings)
	}

	return ratings[len(ratings)-1]
}

func (d *Day) Part2() interface{} {
	counts := d.counts(d.data)
	var oxygenRating []int
	var co2Rating []int

	// Now do a filter
	oxygenRating = d.filterRatings(counts, true)
	co2Rating = d.filterRatings(counts, false)

	return bitutils.Btoi(arrayutils.MapIntToInt8(oxygenRating)) *
		bitutils.Btoi(arrayutils.MapIntToInt8(co2Rating))
}
