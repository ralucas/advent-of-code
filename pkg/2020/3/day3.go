package day3

import (
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type SledState struct {
	Start int
	Pos   int
	End   int
	Right int
	Down  int
}

type Day struct {
	data [][]string
}

func (d *Day) PrepareData(filepath string) {
	data := fileutil.ReadFile(filepath)
	var pData [][]string
	splData := arrayutil.Filter(strings.Split(data, "\n"), func(in string) bool {
		return in != ""
	})
	for _, line := range splData {
		pData = append(pData, strings.Split(line, ""))
	}

	d.data = pData

	return
}

func (d *Day) Part1() interface{} {
	s := SledState{
		Start: 0,
		End:   len(d.data[0]) - 1,
		Right: 3,
		Down:  1,
	}

	treeCount := 0

	for i := 0; i < len(d.data); i += s.Down {
		if s.IsEqualToPosition(d.data[i], "#") {
			treeCount += 1
		}
		s.SetPos(s.NextPosition())
	}

	return treeCount
}

func (d *Day) Part2() interface{} {
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
			Start: 0,
			Right: slope[0],
			Down:  slope[1],
			End:   len(d.data[0]) - 1,
		}

		ss.SetPos(ss.Start)

		for i := 0; i < len(d.data); i += ss.Down {
			if ss.IsEqualToPosition(d.data[i], "#") {
				treeCounts[k] += 1
			}
			ss.SetPos(ss.NextPosition())
		}
		total *= treeCounts[k]
	}

	return total
}

func (s *SledState) IsEqualToPosition(line []string, check string) bool {
	return line[s.Pos] == check
}

func (s *SledState) NextPosition() int {
	nextPos := s.Right + s.Pos
	if nextPos > s.End {
		nextPos = nextPos - s.End - 1
	}

	return nextPos
}

func (s *SledState) SetPos(pos int) {
	s.Pos = pos
}

func (s *SledState) GetPos() int {
	return s.Pos
}
