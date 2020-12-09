package main

import (
	"flag"
	"fmt"

	day9 "github.com/ralucas/advent-of-code/pkg/9"
)

var inputFile = flag.String("input", "assets/9/input.txt", "Input file")

func main() {
	fmt.Print("Day 9\n===========\n")
	flag.Parse()
	data := day9.PrepareData(*inputFile)

	a, aidx := day9.FindFirstNonSum(data, 25)
	fmt.Println("A -- First num:", a)

	b := day9.ContiguousSumSet(data[:aidx], a)
	min, max := day9.Extent(b)
	fmt.Println("B -- Sum of min-max:", min+max)
}
