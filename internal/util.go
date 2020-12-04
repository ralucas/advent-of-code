package util

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFile(filepath string) string {
	f, err := ioutil.ReadFile(filepath); if err != nil {
		log.Fatalf("Error reading in file %v", err)
	}

	return string(f)
}

func ReadFileToArray(filepath string, sep string) []string {
	f, err := ioutil.ReadFile(filepath); if err != nil {
		log.Fatalf("Error reading in file %v", err)
	}

	return Filter(strings.Split(string(f), sep), func(v string) bool {
		return v != ""
	})
}

func MapToInt(vs []string) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		var err error
		vsm[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error processing map to int")
		}
	}
	return vsm
}

func Filter(vs []string, f func(string) bool) []string {
	vsm := make([]string, 0)
	for _, v := range(vs) {
		if f(v) {
			vsm = append(vsm, v)
		}
	}
	return vsm
}

func QSort(vi []int) []int {
	vlen := len(vi)

	if vlen <= 1 {
		return vi
	}

	if vlen == 2 {
		if vi[0] > vi[1] {
			return []int{vi[1], vi[0]}
		}
		return vi
	}

	pivot := vlen / 2

	a, b := make([]int, 0), make([]int, 0)

	// Always remember pivot can't be included in
	// recursive array
	for _, n := range vi {
		if n < vi[pivot] {
			a = append(a, n)
		}
		if n > vi[pivot]  {
			b = append(b, n)
		}
	}

	sorted := make([]int, 0)

	sorted = append(sorted, QSort(a)...)
	sorted = append(sorted, vi[pivot])
	sorted = append(sorted, QSort(b)...)

	return sorted
}
