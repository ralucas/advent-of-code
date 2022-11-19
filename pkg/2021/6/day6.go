package day6

import (
	"log"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data  []int
	Days1 int
	Days2 int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, ",")

	d.data = arrayutil.MapToInt(data)

	return
}

func (d *Day) Part1() interface{} {
	state := NewState(d.data)
	for i := 0; i < d.Days1; i++ {
		state.Day()
	}

	return state.FishCount()
}

func (d *Day) Part2() interface{} {
	buckets := make([]int64, 9)

	for _, d := range d.data {
		buckets[d] += int64(1)
	}

	for i := 0; i < d.Days2; i++ {
		spawns := buckets[0]
		for j := 1; j < len(buckets); j++ {
			buckets[j-1] = buckets[j]
		}
		buckets[6] += spawns
		buckets[len(buckets)-1] = spawns
	}

	count := int64(0)
	for _, b := range buckets {
		count += b
	}

	return count
}
