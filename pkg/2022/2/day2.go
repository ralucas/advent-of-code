package day2

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
)

type Day struct {
	data [][]string
}

type Value int

const (
	Rock     Value = 1
	Paper    Value = 2
	Scissors Value = 3
)

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	d.data = arrayutil.MapTo2D(data, " ")

	return
}

func (d *Day) Part1() interface{} {
	scores := map[string]Value{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	total := 0
	for _, game := range d.data {
		opp, me := game[0], game[1]
		total += int(scores[me])
		total += score(scores[opp], scores[me])
	}

	return total
}

func score(opp, me Value) int {
	switch opp {
	case Rock:
		switch me {
		case Rock:
			return 3
		case Paper:
			return 6
		case Scissors:
			return 0
		}
	case Paper:
		switch me {
		case Rock:
			return 0
		case Paper:
			return 3
		case Scissors:
			return 6
		}
	case Scissors:
		switch me {
		case Rock:
			return 6
		case Paper:
			return 0
		case Scissors:
			return 3
		}
	}

	return 0
}

func (d *Day) Part2() interface{} {
	scores := map[string]Value{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	total := 0

	for _, game := range d.data {
		opp, choice := game[0], game[1]

		switch choice {
		case "X":
			me := loser(scores[opp])
			total += int(me)
		case "Y":
			me := scores[opp]
			total += int(me) + 3
		case "Z":
			me := winner(scores[opp])
			total += int(me) + 6
		}
	}

	return total
}

func loser(v Value) Value {
	switch v {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}

	return 0
}

func winner(v Value) Value {
	switch v {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}

	return 0
}
