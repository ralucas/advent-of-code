package day15

import (
	"log"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data []int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFile(filepath)

	d.data = arrayutil.MapToInt(
		arrayutil.Filter(strings.Split(data, ","), func(s string) bool {
			return s != ""
		}),
	)
}

func (d *Day) Part1() interface{} {
	return FindNumber(d.data, 2020)
}

func (d *Day) Part2() interface{} {
	return FindNumber(d.data, 30000000)
}

func initMap(vi []int) map[int][]int {
	m := make(map[int][]int)

	for i, v := range vi {
		m[v] = []int{i + 1}
	}

	if _, ok := m[0]; ok {
		m[0] = append(m[0], len(vi)+1)
	} else {
		m[0] = []int{len(vi) + 1}
	}

	return m
}

// In this game, the players take turns saying numbers.
// They begin by taking turns reading from a list of starting
// numbers (your puzzle input). Then, each turn consists of
// considering the most recently spoken number:
//
//	If that was the first time the number has been spoken, the current player says 0.
//	Otherwise, the number had been spoken before; the current player announces
//	  how many turns apart the number is from when it was previously spoken.
func FindNumber(vi []int, searchTurn int) int {
	nmap := initMap(vi)
	curNum := 0
	turn := len(vi) + 1

	for {
		if turn == searchTurn {
			return curNum
		}

		if v, ok := nmap[curNum]; ok {
			if len(v) > 1 {
				curNum = v[len(v)-1] - v[len(v)-2]
			} else {
				curNum = turn - v[len(v)-1]
			}
			nmap[curNum] = append(nmap[curNum], turn+1)
		} else {
			curNum = 0
		}

		turn++
	}
}
