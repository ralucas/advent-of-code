package day10

import (
	"log"
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Day struct {
	data          [][]string
	closingPoints map[string]int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, "\n")

	for _, line := range data {
		d.data = append(d.data, strings.Split(line, ""))
	}

	d.closingPoints = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	return
}

func isOpeningBracket(s string) bool {
	return s == "(" || s == "[" || s == "{" || s == "<"
}

func matches(open, close string) bool {
	switch open {
	case "(":
		return close == ")"
	case "[":
		return close == "]"
	case "{":
		return close == "}"
	case "<":
		return close == ">"
	default:
		return false
	}
}

func checkLine(line []string) (string, bool) {
	if !isOpeningBracket(line[0]) {
		return line[0], false
	}

	stack := NewStack(line[0])

	for _, bracket := range line {
		if isOpeningBracket(bracket) {
			stack.Push(bracket)
		} else {
			if !stack.Empty() {
				open, _ := stack.Pop()
				if !matches(open, bracket) {
					return bracket, false
				}
			} else {
				return bracket, false
			}
		}
	}

	return "", true
}

func (d *Day) Part1() interface{} {
	points := make([]int, 0)
	for _, line := range d.data {
		if missing, ok := checkLine(line); !ok {
			points = append(points, d.closingPoints[missing])
		}
	}

	return mathutils.Sum(points)
}

func (d *Day) Part2() interface{} {
	return nil
}
