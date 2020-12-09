package main

import (
	"flag"
	"fmt"

	day5 "github.com/ralucas/advent-of-code/pkg/5"
)

var inputFile = flag.String("input", "", "Input file")

func main() {
	fmt.Print("Day 5\n===========\n")
	flag.Parse()
	data := day5.PrepareData(*inputFile)

	maxID := day5.GetHighestSeatID(data)
	fmt.Println("A -- Max ID:", maxID)

	mySeat := day5.FindSeat(data)
	fmt.Println("B -- Seat ID:", mySeat.ID)
}
