package aoc

import "github.com/ralucas/advent-of-code/pkg/noop"

type AOC interface {
	PrepareData(string)
	Part1() interface{}
	Part2() interface{}
}

func New(day int, year int) AOC {
	switch year {
	case 2020:
		return days2020[day-1]
	case 2021:
		return days2021[day-1]
	case 2022:
		return days2022[day-1]
	default:
		return &noop.Day{}
	}
}
