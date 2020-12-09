package day9

import (
	"log"

	"github.com/ralucas/advent-of-code/pkg/utils"

	day1 "github.com/ralucas/advent-of-code/pkg/1"
)

func PrepareData(filepath string) []int {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	m := utils.MapToInt(data)

	return m
}

func FindFirstNonSum(nums []int, preamble int) (int, int) {
	i := 0
	nlen := len(nums)

	for j := preamble + 1; j < nlen; j++ {
		a, b := day1.TwoSum(nums[i:j], nums[j])
		if a == -1 && b == -1 {
			return nums[j], j
		}
		i++
	}

	return -1, -1
}

func ContiguousSumSet(nums []int, target int) []int {
	nlen := len(nums)

	ch := make(chan []int)

	for i := 0; i < nlen; i++ {
		go func(start int) {
			sumSet := make([]int, 0)
			curSum := nums[start]
			if curSum > target {
				ch <- nil
				return
			}
			sumSet = append(sumSet, nums[start])
			for j := start + 1; j < nlen; j++ {
				curSum += nums[j]
				sumSet = append(sumSet, nums[j])
				if curSum == target {
					ch <- sumSet
					return
				}
				if curSum > target {
					ch <- nil
					return
				}
			}
		}(i)
	}

	for i := 0; i < nlen; i++ {
		ss := <-ch
		if ss != nil {
			return ss
		}
	}

	defer close(ch)
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
