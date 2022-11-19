package day10

import (
	"log"
	"strings"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	data           [][]string
	closingPoints  map[string]int
	closingPoints2 map[string]int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	for _, line := range data {
		d.data = append(d.data, strings.Split(line, ""))
	}

	d.closingPoints = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	d.closingPoints2 = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
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

func closer(open string) string {
	switch open {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	default:
		return ""
	}
}

func checkLine(line []string) (string, bool) {
	if !isOpeningBracket(line[0]) {
		return line[0], false
	}

	stack := NewStack()

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

	return stack.String(), true
}

func (d *Day) Part1() interface{} {
	points := make([]int, 0)
	for _, line := range d.data {
		if missing, ok := checkLine(line); !ok {
			points = append(points, d.closingPoints[missing])
		}
	}

	return mathutil.Sum(points)
}

func (d *Day) Part2() interface{} {
	var closers [][]string

	for _, line := range d.data {
		if left, ok := checkLine(line); ok {
			openers := strings.Split(left, "")
			c := make([]string, 0)
			for _, opener := range openers {
				c = append(c, closer(opener))
			}
			closers = append(closers, c)
		}
	}

	// Start with a total score of 0.
	// Then, for each character, multiply
	// the total score by 5 and then increase
	// the total score by the point value
	// given for the character.
	scores := make([]int, len(closers))

	for i := range scores {
		for j := len(closers[i]) - 1; j >= 0; j-- {
			scores[i] *= 5
			c := closers[i][j]
			scores[i] += d.closingPoints2[c]
		}
	}

	sorted := sortutil.QSort(scores)

	return mathutil.Median(sorted)
}
