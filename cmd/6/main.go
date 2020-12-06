package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	util "github.com/ralucas/advent-of-code/internal"
)

var inputFile = flag.String("input", "", "Input file")

func prepareData(filepath string) [][]string {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := util.ReadFileToArray(filepath, "\n\n")

	prepared := make([][]string, len(data))

	for i, d := range data {
		prepared[i] = strings.Split(d, "\n")
	}

	return prepared
}

func groupCount(vs []string) int {
	alphaMap := make(map[int32]bool)

	count := 0

	for _, v := range vs {
		for _, c := range v {
			if _, ok := alphaMap[c]; !ok {
				alphaMap[c] = true
				count += 1
			}
		}
	}

	return count
}

func allYesCount(vs []string) int {
	alphaMap := make(map[int32]int)

	count := 0

	grouplen := len(vs)

	for _, v := range vs {
		for _, c := range v {
			alphaMap[c] += 1
			if alphaMap[c] == grouplen {
				count += 1
			}
		}
	}

	return count
}

func sumCounts(vvs [][]string, fn func([]string) int) int {
	sum := 0
	for _, vs := range vvs {
		gc := fn(vs)
		sum += gc
	}

	return sum
}

func main() {
	fmt.Print("Day 6\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	sumPart1 := sumCounts(data, groupCount)
	fmt.Println("A -- Group Count Sum:", sumPart1)

	sumPart2 := sumCounts(data, allYesCount)
	fmt.Println("B -- All Yes Count Sum:", sumPart2)
}
