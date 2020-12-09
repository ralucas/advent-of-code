package main

import (
	"fmt"

	day2 "github.com/ralucas/advent-of-code/pkg/2"
)

func main() {
	data := day2.PrepareData("assets/2/input.txt")

	validCount := 0
	validPosCount := 0

	for _, pw := range data {
		if day2.IsValid(pw) {
			validCount += 1
		}
		if day2.IsValidByPosition(pw) {
			validPosCount += 1
		}
	}

	fmt.Println("Valid Count:", validCount)
	fmt.Println("Valid Positional Count:", validPosCount)
}
