package day6

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	"github.com/ralucas/advent-of-code/pkg/util/queue"
	// arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	// bitutil "github.com/ralucas/advent-of-code/pkg/util/bit"
	// mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	// sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	data string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}

	d.data = fileutil.ReadFile(filepath)

	return
}

func put(m map[byte]int, val byte) map[byte]int {
	if _, ok := m[val]; !ok {
		m[val] = 1
	} else {
		m[val] += 1
	}

	return m
}

// sliding window
func FirstUniqueN(s string, n int) int {
	// instantiate and fill queue with 1st 4
	q := queue.New([]byte(s[:n])...)

	// instantiate and fill map with 1st n
	m := make(map[byte]int)
	for i := 0; i < n; i++ {
		val := s[i]

		m = put(m, val)

		if len(m) == n {
			return i + 1
		}
	}

	for i := n; i < len(s); i++ {
		k, err := q.Pop()
		if err != nil {
			log.Fatal(err)
		}

		m[k] -= 1
		if m[k] == 0 {
			delete(m, k)
		}

		val := s[i]

		q.Push(val)

		m = put(m, val)

		if len(m) == n {
			return i + 1
		}
	}

	return -1
}

func (d *Day) Part1() interface{} {
	return FirstUniqueN(d.data, 4)
}

func (d *Day) Part2() interface{} {
	return FirstUniqueN(d.data, 14)
}
