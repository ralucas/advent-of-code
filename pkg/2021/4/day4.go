package day4

import (
	"log"
	"regexp"
	"strings"

	"github.com/ralucas/advent-of-code/pkg/utils"
	arrayutil "github.com/ralucas/advent-of-code/pkg/utils/array"
)

type Day struct {
	// TODO: Change this
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
	data := utils.ReadFileToArray(filepath, "\n")

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
	return nil
}

func (d *Day) Part2() interface{} {
	return nil
}
