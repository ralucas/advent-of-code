package day13

import (
	"log"
	"strconv"
	"strings"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	earliestTimestamp int
	busSchedule       []string
	buses             []int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	var err error
	d.earliestTimestamp, err = strconv.Atoi(data[0])
	if err != nil {
		log.Fatalf("Failed to parse earliest timestamp")
	}

	d.busSchedule = strings.Split(data[1], ",")

	filtered := utils.Filter(d.busSchedule, func(s string) bool {
		return s != "x"
	})

	d.buses = utils.MapToInt(filtered)

	return
}

func (d *Day) Part1() interface{} {
	min, minBus := NearestNextBus(d.earliestTimestamp, d.buses)

	return min * minBus
}

func (d *Day) Part2() interface{} {
	return -1
}

func NearestNextBus(ts int, buses []int) (int, int) {
	min := int(^uint(0) >> 1)
	minBus := buses[0]

	for _, bus := range buses {
		div := ts / bus
		next := (div + 1) * bus
		diff := next - ts
		if diff < min {
			min = diff
			minBus = bus
		}
	}

	return min, minBus
}
