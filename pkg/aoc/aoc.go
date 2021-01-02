package aoc

import (
	day1 "github.com/ralucas/advent-of-code/pkg/1"
	day10 "github.com/ralucas/advent-of-code/pkg/10"
	day11 "github.com/ralucas/advent-of-code/pkg/11"
	day12 "github.com/ralucas/advent-of-code/pkg/12"
	day13 "github.com/ralucas/advent-of-code/pkg/13"
	day14 "github.com/ralucas/advent-of-code/pkg/14"
	day15 "github.com/ralucas/advent-of-code/pkg/15"
	day16 "github.com/ralucas/advent-of-code/pkg/16"
	day2 "github.com/ralucas/advent-of-code/pkg/2"
	day3 "github.com/ralucas/advent-of-code/pkg/3"
	day4 "github.com/ralucas/advent-of-code/pkg/4"
	day5 "github.com/ralucas/advent-of-code/pkg/5"
	day6 "github.com/ralucas/advent-of-code/pkg/6"
	day7 "github.com/ralucas/advent-of-code/pkg/7"
	day8 "github.com/ralucas/advent-of-code/pkg/8"
	day9 "github.com/ralucas/advent-of-code/pkg/9"
	"github.com/ralucas/advent-of-code/pkg/noop"
)

type AOC interface {
	PrepareData(string)
	Part1() interface{}
	Part2() interface{}
}

func New(day int) AOC {
	switch day {
	case 1:
		return &day1.Day{}
	case 2:
		return &day2.Day{}
	case 3:
		return &day3.Day{}
	case 4:
		return &day4.Day{}
	case 5:
		return &day5.Day{}
	case 6:
		return &day6.Day{}
	case 7:
		return &day7.Day{}
	case 8:
		return &day8.Day{}
	case 9:
		return &day9.Day{}
	case 10:
		return &day10.Day{}
	case 11:
		return &day11.Day{}
	case 12:
		return &day12.Day{}
	case 13:
		return &day13.Day{}
	case 14:
		return &day14.Day{}
	case 15:
		return &day15.Day{}
	case 16:
		return &day16.Day{}
	//case DAYX:
	//    return &dayDAYX.Day{}
	default:
		return &noop.Day{}
	}
}
