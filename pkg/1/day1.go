package day1

import (
	"strings"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

func PrepareData(filepath string) []int {
	inputString := utils.ReadFile(filepath)
	inputArr := utils.Filter(strings.Split(inputString, "\n"), func(s string) bool {
		return s != ""
	})

	preparedData := utils.MapToInt(inputArr)

	return preparedData
}

// TwoSum takes an unsorted array and finds
// the two numbers that add up to the target.
// Runtime is O(n), space O(n).
func TwoSum(data []int, target int) (int, int) {
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

// TwoSumSorted takes an sorted array and finds
// the two numbers that add up to the target.
// Runtime is O(n), space O(1).
func TwoSumSorted(sData []int, target int) (int, int) {
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

// ThreeSum takes array and target and returns the
// first three numbers that add up to target.
// Requires input array data to be sorted.
// Runtime O(n^2), space O(1)
func ThreeSum(data []int, target int) (int, int, int) {
	for i := 0; i < len(data); i++ {
		b, c := TwoSumSorted(data[i+1:], target-data[i])
		add3 := data[i] + b + c
		if add3 == target {
			return data[i], b, c
		}
	}

	return -1, -1, -1
}
