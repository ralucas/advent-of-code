package main

import (
	"fmt"

	day3 "github.com/ralucas/advent-of-code/pkg/3"
)

func main() {
	fmt.Println("Running AOC #3...")

	data := day3.PrepareData("assets/3/input.txt")

	s := day3.SledState{
		Start: 0,
		End:   len(data[0]) - 1,
		Right: 3,
		Down:  1,
	}

	treeCount := 0

	for i := 0; i < len(data); i += s.Down {
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
		ss := day3.SledState{
			Start: 0,
			Right: slope[0],
			Down:  slope[1],
			End:   len(data[0]) - 1,
		}

		ss.SetPos(ss.Start)

		for i := 0; i < len(data); i += ss.Down {
			if ss.IsEqualToPosition(data[i], "#") {
				treeCounts[k] += 1
			}
			ss.SetPos(ss.NextPosition())
		}
		total *= treeCounts[k]
	}

	fmt.Println("B -- Tree Multiples:", total)
}
