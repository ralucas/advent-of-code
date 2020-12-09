package main

import (
	"flag"
	"fmt"

	day6 "github.com/ralucas/advent-of-code/pkg/6"
)

var inputFile = flag.String("input", "", "Input file")

func main() {
	fmt.Print("Day 6\n===========\n")
	flag.Parse()
	data := day6.PrepareData(*inputFile)

	sumPart1 := day6.SumCounts(data, day6.GroupCount)
	fmt.Println("A -- Group Count Sum:", sumPart1)

	sumPart2 := day6.SumCounts(data, day6.AllYesCount)
	fmt.Println("B -- All Yes Count Sum:", sumPart2)
}
