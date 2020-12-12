package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ralucas/advent-of-code/pkg/aoc"
)

var inputFile = flag.String("input", "", "Input file")
var day = flag.Int("day", getDay(), "Day to run")

func getDay() int {
	return time.Now().Day()
}

func run(runner aoc.AOC, inputFile string) {
	runner.PrepareData(inputFile)
	part1 := runner.Part1()
	part2 := runner.Part2()

	fmt.Println("  Part 1:", part1)
	fmt.Println("  Part 2:", part2)
}

func main() {
	flag.Parse()

	fmt.Printf("\n\n================\n    Day %d\n================\n\n", *day)

	run(aoc.New(*day), *inputFile)

	fmt.Print("\n================\n")
}
