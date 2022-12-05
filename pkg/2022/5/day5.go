package day5

import (
	"log"
	"strconv"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	stackData []string
	moves     [][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}

	data := fileutil.ReadFileToArray(filepath, "\n\n")
	d.stackData = strings.Split(data[0], "\n")
	movesCsv := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(data[1], "move ", ""), " from ", ","), " to ", ",")
	moves := strings.Split(movesCsv, "\n")
	d.moves = arrayutil.MapTo2DGen(moves, func(s string, _ int) []int {
		ss := strings.Split(s, ",")
		return arrayutil.MapToInt(ss)
	})

	return
}

func (d *Day) buildStacks() []*Stack {
	numLine := strings.TrimSpace(d.stackData[len(d.stackData)-1])
	lens, err := strconv.Atoi(string(numLine[len(numLine)-1]))
	if err != nil {
		log.Fatalf("%v", err)
	}

	stacks := make([]*Stack, lens)
	for i := range stacks {
		stacks[i] = NewStack()
	}

	data0 := d.stackData[:len(d.stackData)-1]

	for x := len(data0) - 1; x >= 0; x -= 1 {
		row := data0[x]
		rlen := len(row)

		idx := 0
		for i := 1; i < rlen; i += 4 {
			item := strings.TrimSpace(string(row[i]))
			if item != "" {
				stacks[idx].Push(item)
			}
			idx += 1
		}
	}

	return stacks
}

func (d *Day) Part1() interface{} {
	stacks := d.buildStacks()

	for _, move := range d.moves {
		qty, from, to := move[0], move[1]-1, move[2]-1

		for i := 0; i < qty; i++ {
			moveVal, err := stacks[from].Pop()
			if err != nil {
				log.Fatal(err)
			}
			stacks[to].Push(moveVal)
		}
	}

	output := ""

	for _, stack := range stacks {
		output += stack.Peek()
	}

	return output
}

func (d *Day) Part2() interface{} {
	stacks := d.buildStacks()

	for _, move := range d.moves {
		qty, from, to := move[0], move[1]-1, move[2]-1

		moveVals, err := stacks[from].PopN(qty)
		if err != nil {
			log.Fatalf("stack #%d %+v", from, err)
		}
		stacks[to].Push(moveVals...)
	}

	output := ""

	for _, stack := range stacks {
		output += stack.Peek()
	}

	return output
}
