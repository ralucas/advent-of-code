package array

import (
	"log"
	"strconv"
	"strings"
)

func MapTo2D(vs []string, sep string) [][]string {
	vsm := make([][]string, len(vs))
	for i, v := range vs {
		vsm[i] = strings.Split(v, sep)
	}

	return vsm
}

func MapToInt(vs []string) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		var err error
		vsm[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error processing map to int %v\n", err)
		}
	}
	return vsm
}

func MapToInt8(vs []string) []int8 {
	vsm := make([]int8, len(vs))
	for i, v := range vs {
		var err error
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error processing map to int %v\n", err)
		}

		vsm[i] = int8(val)
	}

	return vsm
}

func MapIntToInt8(vi []int) []int8 {
	vsm := make([]int8, len(vi))
	for i, v := range vi {
		vsm[i] = int8(v)
	}

	return vsm
}

func Filter(vs []string, f func(string) bool) []string {
	vsm := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsm = append(vsm, v)
		}
	}
	return vsm
}

func FilterInt(vi []int, f func(int) bool) []int {
	vim := make([]int, 0)
	for _, v := range vi {
		if f(v) {
			vim = append(vim, v)
		}
	}

	return vim
}

func FilterInt2D(vi [][]int, f func([]int) bool) [][]int {
	vim := make([][]int, 0)
	for _, v := range vi {
		if f(v) {
			vim = append(vim, v)
		}
	}

	return vim
}

func FindIntIndexes(vi []int, f func(int) bool) []int {
	vim := make([]int, 0)

	for i, v := range vi {
		if f(v) {
			vim = append(vim, i)
		}
	}

	return vim
}

func Every(vi []int, f func(int) bool) bool {
	out := true

	for _, v := range vi {
		if !f(v) {
			return false
		}
	}

	return out
}

// Index returns the first index found
// where the input value is found,
// else -1 if nothing found.
func Index(vs []string, val string) int {
	for i, v := range vs {
		if v == val {
			return i
		}
	}

	return -1
}

func IndexInt(vi []int, val int) int {
	for i, v := range vi {
		if v == val {
			return i
		}
	}

	return -1
}

// IndexesInt returns all indexes that match
func IndexesInt(vi []int, val int) []int {
	var output []int

	for i, v := range vi {
		if v == val {
			output = append(output, i)
		}
	}

	return output
}

func Equal(v1, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}

	l, r := 0, len(v1)-1

	for l <= r {
		if v1[l] != v2[l] {
			return false
		}
		if v1[r] != v2[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func EqualDistance(v1, v2 []int) int {
	if len(v1) != len(v2) {
		return -1
	}

	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			return len(v1) - i
		}
	}

	return 0
}
