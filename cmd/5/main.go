package main

import (
	"flag"
	"fmt"
	util "github.com/ralucas/advent-of-code/internal"
	"log"
	"strings"
)

var inputFile = flag.String("input", "", "Input file")

type BoardingPass struct {
	row int
	col int
	id int
}

func prepareData(filepath string) [][]string {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := util.ReadFileToArray(filepath, "\n")

	prepared := make([][]string, len(data))

	for i, d := range data {
		prepared[i] = strings.Split(d, "")
	}

	return prepared
}

func getRow(s []string) int {
	return getNum(0, 127, "F", "B", s)
}

func getCol(s []string) int {
	return getNum(0, 7, "L", "R", s)
}

func getNum(min, max int, minL, maxL string, s []string) int {
	for _, c := range s[:len(s)-1] {
		if c != minL && c != maxL {
			return -1
		}
		switch c {
		case minL:
			max = (min + max) / 2
		case maxL:
			min = ((min + max) / 2) + 1
		}
	}

	if s[len(s)-1] == minL {
		return min
	}

	return max
}

// toBoardingPass converts a boarding pass string to numeric values.
// Every seat also has a unique seat ID: multiply the row by 8, then add the column
func toBoardingPass(s []string) BoardingPass {
	row := getRow(s[:7])
	col := getCol(s[7:])

	bp := BoardingPass{
		row: row,
		col: col,
		id: (row * 8) + col,
	}

	return bp
}

func getHighestSeatID(ss [][]string) int {
	var maxID int

	for _, s := range(ss) {
		bp := toBoardingPass(s)
		if bp.id > maxID {
			maxID = bp.id
		}
	}

	return maxID
}

func findSeat(ss [][]string) BoardingPass {
	plane := make([][]int, 127)
	for i := range plane {
		plane[i] = make([]int, 7)
	}

	rows := make([]int, 127)

	for _, s := range ss {
		bp := toBoardingPass(s)
		plane[bp.row][bp.col] = 1
	}

	// TODO: Check for empty seats
	for i, row := range plane {
		if util.Every(row, func(v int) bool { return v == 1 }) {
			rows[i] = 1
		}
	}
}

func main() {
	fmt.Print("Day 5\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	maxID := getHighestSeatID(data)
	fmt.Println("A -- Max ID:", maxID)
}

