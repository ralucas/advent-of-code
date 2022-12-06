package day6_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day6 "github.com/ralucas/advent-of-code/pkg/2022/6"
)

var td day6.Day

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expect: 7},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expect: 5},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", expect: 6},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expect: 10},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expect: 11},
	}
	for _, tc := range tests {
		result := day6.FirstUniqueN(tc.input, 4)
		assert.Equal(t, tc.expect, result)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expect: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expect: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", expect: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expect: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expect: 26},
	}
	for _, tc := range tests {
		result := day6.FirstUniqueN(tc.input, 14)
		assert.Equal(t, tc.expect, result)
	}
}
