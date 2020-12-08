package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	util "github.com/ralucas/advent-of-code/internal"
)

func prepareData(filepath string) []int {
	inputString := util.ReadFile(filepath)
	inputArr := util.Filter(strings.Split(inputString, "\n"), func(s string) bool {
		return s != ""
	})

	preparedData := util.MapToInt(inputArr)

	return preparedData
}

// twoSum takes an unsorted array and finds
// the two numbers that add up to the target.
// Runtime is O(n), space O(n).
func twoSum(data []int, target int) (int, int) {
	i, j := 0, len(data)-1
	m := make(map[int]int)

	for i < j {
		ti := target - data[i]
		if _, ok := m[ti]; ok {
			return ti, data[i]
		}
		m[data[i]] = i
		tj := target - data[j]
		if _, ok := m[tj]; ok {
			return tj, data[j]
		}
		m[data[j]] = j
		i++
		j--
	}

	return -1, -1
}

// twoSumSorted takes an sorted array and finds
// the two numbers that add up to the target.
// Runtime is O(n), space O(1).
func twoSumSorted(sData []int, target int) (int, int) {
	l, r := 0, len(sData)-1

	for l < r {
		add2 := sData[l] + sData[r]
		if add2 == target {
			return sData[l], sData[r]
		}
		if add2 > target {
			r--
		} else {
			l++
		}
	}
	return -1, -1
}

// threeSum takes array and target and returns the
// first three numbers that add up to target.
// Requires input array data to be sorted.
// Runtime O(n^2), space O(1)
func threeSum(data []int, target int) (int, int, int) {
	for i := 0; i < len(data); i++ {
		b, c := twoSumSorted(data[i+1:], target-data[i])
		add3 := data[i] + b + c
		if add3 == target {
			return data[i], b, c
		}
	}

	return -1, -1, -1
}

func main() {
	data := prepareData("assets/1/input.txt")

	target := flag.Int("target", 2020, "target")

	a, b := twoSum(data, *target)
	if a == -1 && b == -1 {
		log.Fatalf("Couldn't find entries")
	}
	result := a * b
	fmt.Println("Two Entry Result: ", result)

	sData := util.QSort(data)
	c, d, e := threeSum(sData, *target)
	if c == -1 {
		log.Fatalf("Couldn't find entries")
	}
	fmt.Println(c, d, e)
	result2 := c * d * e
	fmt.Println("Three Entry Result: ", result2)
}
