package main

import (
	"flag"
	"fmt"

	day7 "github.com/ralucas/advent-of-code/pkg/7"
)

var inputFile = flag.String("input", "", "Input file")

func main() {
	fmt.Print("Day 7\n===========\n")
	flag.Parse()
	data := day7.PrepareData(*inputFile)

	count, _ := day7.CountParents("shiny_gold", data)

	fmt.Println("A -- Count:", count)

	bcount, _ := day7.CountContains("shiny_gold", data)

	fmt.Println("B -- Count:", bcount)
}
