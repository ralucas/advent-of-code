package main

import (
	"flag"
	"fmt"
	"log"

	one "github.com/ralucas/advent-of-code/cmd/1"
	util "github.com/ralucas/advent-of-code/internal"
)

var inputFile = flag.String("input", "assets/9/input.txt", "Input file")

func prepareData(filepath string) []int {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := util.ReadFileToArray(filepath, "\n")

	m := util.MapToInt(data)

	return m
}

func FindFirstNonSum(nums []int, preamble int) (int, int) {
	i := 0
	nlen := len(nums)
	for j := preamble + 1; j < nlen; j++ {
		a, b := one.TwoSum(nums[i:j], nums[j])
		if a == -1 && b == -1 {
			return nums[j], j
		}
		i++
	}

	return -1, -1
}

func ContiguousSumSet(nums []int, target int) []int {
	nlen := len(nums)

	for i := 0; i < nlen; i++ {
		sumSet := make([]int, 0)
		curSum := nums[i]
		if curSum > target {
			continue
		}
		sumSet = append(sumSet, nums[i])
		for j := i + 1; j < nlen; j++ {
			curSum += nums[j]
			sumSet = append(sumSet, nums[j])
			if curSum == target {
				return sumSet
			}
			if curSum > target {
				break
			}
		}
	}

	return nil
}

func MinMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	fmt.Print("Day 9\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	a, aidx := FindFirstNonSum(data, 25)
	fmt.Println("A -- First num:", a)

	b := ContiguousSumSet(data[:aidx], a)
	min, max := MinMax(b)
	fmt.Println("B -- Sum of min-max:", min+max)
}
