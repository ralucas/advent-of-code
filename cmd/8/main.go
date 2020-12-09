package main

import (
	"flag"
	"fmt"
	"log"

	day8 "github.com/ralucas/advent-of-code/pkg/8"
)

var inputFile = flag.String("input", "assets/8/input.txt", "Input file")

func main() {
	fmt.Print("Day 8\n===========\n")
	flag.Parse()
	data := day8.PrepareData(*inputFile)

	lastAcc, exitcode := day8.RunInstructions(data)
	if exitcode != -1 {
		log.Fatalf("Error, all instructions ran")
	}
	fmt.Println("A -- Last Accumulator:", lastAcc)

	finalAcc := day8.FixInstructions(data)
	fmt.Println("B -- Final Accumulator:", finalAcc)
}
