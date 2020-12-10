package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ralucas/advent-of-code/pkg/aoc"
)

var inputFile = flag.String("input", "", "Input file")
var day = flag.Int("day", getDay(), "Day to run")

func run(day int, inputFile string) {
	runner := aoc.New(day)

	runner.PrepareData(inputFile)
	part1 := runner.Part1()
	part2 := runner.Part2()

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func getDay() int {
	return time.Now().Day()
}

func main() {
	flag.Parse()

	fmt.Printf("\nDay %d\n===========\n\n", *day)

	run(*day, *inputFile)

	fmt.Print("\n================\n")
}
