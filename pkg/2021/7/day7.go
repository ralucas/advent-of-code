package day7

import (
	"log"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	data []int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, ",")

	d.data = arrayutil.MapToInt(data)

	return
}

// min = |m - a1| + |m - a2| ...
// [0 1 1 2 4 7 14 16]
// [2 1 1 0 -2 -5 -12 -14]
// That's not the right answer; your answer is too low.
// If you're stuck, make sure you're using the full input data;
// there are also some general tips on the about page, or you can ask for hints on the subreddit.
// Please wait one minute before trying again. (You guessed 268285.) [Return to Day 7]
func (d *Day) Part1() interface{} {
	sorted := sortutil.QSort(d.data)
	min, max := sorted[0], sorted[len(sorted)-1]

	ans := []int{mathutil.MaxInt, max}
	for i := min; i < max; i++ {
		test := 0
		m := i
		for _, v := range d.data {
			test += mathutil.Abs(m - v)
		}
		if test < ans[0] {
			ans[0] = test
			ans[1] = m
		}
	}

	return ans[0]
}

func (d *Day) Part2() interface{} {
	sorted := sortutil.QSort(d.data)
	min, max := sorted[0], sorted[len(sorted)-1]

	ans := []int{mathutil.MaxInt, max}

	for i := min; i < max; i++ {
		test := 0
		m := i
		for _, v := range d.data {
			test += calculateFuelCost(mathutil.Abs(m - v))
		}
		if test < ans[0] {
			ans[0] = test
			ans[1] = m
		}
	}

	return ans[0]
}

func calculateFuelCost(n int) int {
	cost := 0
	for i := 1; i < n+1; i++ {
		cost += i
	}

	return cost
}
