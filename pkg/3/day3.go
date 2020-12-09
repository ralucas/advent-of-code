package day3

import (
	"strings"

	"github.com/ralucas/advent-of-code/pkg/util"
)

type SledState struct {
	Start int
	Pos   int
	End   int
	Right int
	Down  int
}

func PrepareData(filepath string) [][]string {
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
