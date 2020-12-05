package main

import (
	"flag"
	"fmt"
	util "github.com/ralucas/advent-of-code/internal"
	"log"
)

var inputFile = flag.String("input", "", "Input file")

// TODO: Alter this for actual implementation
func prepareData(filepath string) []string {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := util.ReadFileToArray(filepath, "\n")

	return data
}

func main() {
	fmt.Print("Day {Day # HERE}\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	fmt.Printf("%+v", data)
}

