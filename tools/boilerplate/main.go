package main

import (
	"flag"
	"fmt"
	"log"

	util "github.com/ralucas/advent-of-code/internal"
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
	fmt.Print("Day %%DAY%%\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	fmt.Printf("%+v", data)
}
