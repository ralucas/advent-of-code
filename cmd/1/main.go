package main

import (
	"flag"
	"fmt"
	"log"

	day1 "github.com/ralucas/advent-of-code/pkg/1"

	sort "github.com/ralucas/advent-of-code/pkg/utils"
)

func main() {
	data := day1.PrepareData("assets/1/input.txt")

	target := flag.Int("target", 2020, "target")

	a, b := day1.TwoSum(data, *target)
	if a == -1 && b == -1 {
		log.Fatalf("Couldn't find entries")
	}
	result := a * b
	fmt.Println("Two Entry Result: ", result)

	sData := sort.QSort(data)
	c, d, e := day1.ThreeSum(sData, *target)
	if c == -1 {
		log.Fatalf("Couldn't find entries")
	}
	fmt.Println(c, d, e)
	result2 := c * d * e
	fmt.Println("Three Entry Result: ", result2)
}
