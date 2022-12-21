package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ralucas/advent-of-code/pkg/aoc"
)

var (
	day       = flag.Int("day", getDay(), "Day to run")
	year      = flag.Int("year", getYear(), "Year to run")
	inputFile = flag.String("input", getFile(), "Input file")
)

func getDay() int {
	loc, _ := time.LoadLocation("America/New_York")
	return time.Now().In(loc).Day()
}

func getYear() int {
	loc, _ := time.LoadLocation("America/New_York")
	return time.Now().In(loc).Year()
}

func getFile() string {
	return fmt.Sprintf("../assets/%d/%d/input.txt", *year, *day)
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

	fmt.Printf("\n\n================\n  Day %d - %d\n================\n\n", *day, *year)

	run(aoc.New(*day, *year), *inputFile)

	fmt.Print("\n================\n")
}
