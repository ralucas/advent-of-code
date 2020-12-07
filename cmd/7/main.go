package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	util "github.com/ralucas/advent-of-code/internal"
)

var inputFile = flag.String("input", "", "Input file")

func prepareData(filepath string) map[string]map[string]int {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := util.ReadFileToArray(filepath, "\n")

	output := make(map[string]map[string]int)

	for _, line := range data {
		fields := strings.Split(line, "contain")
		k := strings.TrimSpace(strings.Split(fields[0], "bag")[0])
		key := strings.ReplaceAll(k, " ", "_")

		if strings.TrimSpace(fields[1]) == "no other bags." {
			output[key] = nil
			continue
		}
		vals := strings.Split(fields[1], ",")
		valMap := make(map[string]int)
		reQty := regexp.MustCompile(`\d+`)
		reColor := regexp.MustCompile(`\d+\s([A-Za-z\s]+)(\sbag)`)
		for _, v := range vals {
			qty, err := strconv.Atoi(reQty.FindString(v))
			if err != nil {
				log.Fatalf("Error converting str to num %s", v)
			}
			matches := reColor.FindStringSubmatch(v)
			valKey := strings.ReplaceAll(matches[1], " ", "_")
			valMap[valKey] = qty
		}
		output[key] = valMap
	}

	return output
}

func countParents(needle string, haystack map[string]map[string]int) (int, map[string]int) {
	count := 0
	parents := make(map[string]int)

	for k, v := range haystack {
		if _, ok := v[needle]; ok {
			count += 1
			parents[k] += 1
		}
	}

	for k := range parents {
		curCount, curParents := countParents(k, haystack)
		for cp := range curParents {
			if _, ok := parents[cp]; ok {
				parents[cp] += 1
				curCount -= 1
			} else {
				parents[cp] += 1
			}
		}
		count += curCount
	}

	return count, parents
}

func countContains(needle string, haystack map[string]map[string]int) (int, map[string]int) {
	count := 0
	contains := make(map[string]int)

	valMap := haystack[needle]

	for k, v := range valMap {
		contains[k] += v
		_, curContains := countContains(k, haystack)
		for ck, cv := range curContains {
			contains[ck] += cv * v
		}
	}

	for _, v := range contains {
		count += v
	}

	return count, contains
}

func main() {
	fmt.Print("Day 7\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	count, _ := countParents("shiny_gold", data)

	fmt.Println("A -- Count:", count)

	bcount, _ := countContains("shiny_gold", data)

	fmt.Println("B -- Count:", bcount)
}
