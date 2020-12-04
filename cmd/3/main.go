package main

import (
	"fmt"
	util "github.com/ralucas/advent-of-code/internal"
	"strings"
)

type SledState struct {
	start int
	pos int
	end int
	right int
	down int
}

func prepareData(filepath string) [][]string {
	data := util.ReadFile(filepath)
	var pData [][]string
	splData := util.Filter(strings.Split(data, "\n"), func(in string) bool {
		return in != ""
	})
	for _, line := range splData {
		pData = append(pData, strings.Split(line, ""))
	}

	return pData
}

func (s *SledState) IsEqualToPosition(line []string, check string) bool {
	return line[s.pos] == check
}

func (s *SledState) NextPosition() int {
	nextPos := s.right + s.pos
	if nextPos > s.end {
		nextPos = nextPos - s.end - 1
	}

	return nextPos
}

func (s *SledState) SetPos(pos int) {
	s.pos = pos
}

func (s *SledState) GetPos() int {
	return s.pos
}

func main() {
	fmt.Println("Running AOC #3...")

	data := prepareData("assets/3/input.txt")

	s := SledState{
		start: 0,
		end: len(data[0]) - 1,
		right: 3,
		down: 1,
	}

	s.SetPos(s.start)

	treeCount := 0

	for i := 0; i < len(data); i += s.down {
		if s.IsEqualToPosition(data[i], "#") {
			treeCount += 1
		}
		s.SetPos(s.NextPosition())
	}

	fmt.Println("A -- Tree Count:", treeCount)

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeCounts := make([]int, 5)
	total := 1

	for k, slope := range slopes {
		ss := SledState{
			start: 0,
			right: slope[0],
			down: slope[1],
			end: len(data[0]) - 1,
		}

		ss.SetPos(ss.start)

		for i := 0; i < len(data); i += ss.down {
			if ss.IsEqualToPosition(data[i], "#") {
				treeCounts[k] += 1
			}
			ss.SetPos(ss.NextPosition())
		}
		total *= treeCounts[k]
	}

	fmt.Println("B -- Tree Multiples:", total)
}