package day6

import (
	"log"
	"strings"

	"github.com/ralucas/advent-of-code/pkg/util"
)

func PrepareData(filepath string) [][]string {
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

func GroupCount(vs []string) int {
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

func AllYesCount(vs []string) int {
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

func SumCounts(vvs [][]string, fn func([]string) int) int {
	sum := 0
	for _, vs := range vvs {
		gc := fn(vs)
		sum += gc
	}

	return sum
}
