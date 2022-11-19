package day4

import (
	"log"
	"regexp"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
)

type Day struct {
	numbers []int
	boards  []*Board
}

func (d *Day) Numbers() []int {
	return d.numbers
}

func (d *Day) Boards() []*Board {
	return d.boards
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	d.numbers = arrayutil.MapToInt(strings.Split(data[0], ","))

	d.boards = make([]*Board, 0)

	boardVals := make([][]int, 5)

	re := regexp.MustCompile(`\s+`)
	for i, line := range data[1:] {
		boardVals[i%5] = make([]int, 5)
		tl := strings.TrimSpace(line)
		pl := re.ReplaceAll([]byte(tl), []byte(","))
		boardVals[i%5] = arrayutil.MapToInt(strings.Split(string(pl), ","))

		if boardVals[4] != nil {
			d.boards = append(d.boards, NewBoard(boardVals))
			boardVals = make([][]int, 5)
		}
	}

	return
}

func (d *Day) Part1() interface{} {
	for _, n := range d.numbers {
		for _, board := range d.boards {
			bingo := board.Mark(n)
			if bingo {
				unmarked := board.UnmarkedValues()
				s := mathutil.Sum(unmarked)
				return s * n
			}
		}
	}

	return nil
}

func (d *Day) Part2() interface{} {
	boards := make([]*Board, len(d.boards))
	copy(boards, d.boards)

	nilCount := 0

	for _, n := range d.numbers {
		for i, board := range boards {
			if board != nil {
				bingo := board.Mark(n)
				if bingo && len(boards)-nilCount == 1 {
					unmarked := board.UnmarkedValues()
					s := mathutil.Sum(unmarked)
					return s * n
				}
				if bingo {
					boards[i] = nil
					nilCount += 1
				}
			}
		}
	}

	return nil
}
