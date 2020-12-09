package main

import (
	"flag"
	"fmt"

	day4 "github.com/ralucas/advent-of-code/pkg/4"
)

var inputFile = flag.String("input", "", "Input file")

func main() {
	flag.Parse()
	data := day4.PrepareData(*inputFile)

	validPassportCount := day4.CountValidPassports(data)
	fmt.Println("A -- Valid Passport Count:", validPassportCount)

	validPassportStrictCount := day4.CountValidPassportsStrict(data)
	fmt.Println("B -- Valid Passport Strict Count:", validPassportStrictCount)
}
