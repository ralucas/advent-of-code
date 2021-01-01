package day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint
var td Day

func init() {
	td.PrepareData("../../test/testdata/15/test_input.txt")
}

func TestFindNumberSteps(t *testing.T) {
	tests := []struct {
		expect int
		step   int
	}{
		{expect: 0, step: 4},
		{expect: 3, step: 5},
		{expect: 3, step: 6},
		{expect: 1, step: 7},
		{expect: 0, step: 8},
		{expect: 4, step: 9},
		{expect: 0, step: 10},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%dStep%d", i, test.step), func(t *testing.T) {
			assert.Equal(t, test.expect, FindNumber([]int{0, 3, 6}, test.step))
		})
	}
}

func TestFindNumber(t *testing.T) {
	t.Run("Find2020", func(t *testing.T) {
		tests := []struct {
			expect int
			input  []int
		}{
			{input: []int{0, 3, 6}, expect: 436},
			{input: []int{1, 3, 2}, expect: 1},
			{input: []int{2, 1, 3}, expect: 10},
			{input: []int{1, 2, 3}, expect: 27},
			{input: []int{2, 3, 1}, expect: 78},
			{input: []int{3, 2, 1}, expect: 438},
			{input: []int{3, 1, 2}, expect: 1836},
		}

		for i, test := range tests {
			t.Run(fmt.Sprintf("Test%dExpect%d", i, test.expect), func(t *testing.T) {
				assert.Equal(t, test.expect, FindNumber(test.input, 2020))
			})
		}
	})

	t.Run("Find30000000th", func(t *testing.T) {
		tests := []struct {
			expect int
			input  []int
		}{
			{input: []int{0, 3, 6}, expect: 175594},
			{input: []int{1, 3, 2}, expect: 2578},
			{input: []int{2, 1, 3}, expect: 3544142},
			{input: []int{1, 2, 3}, expect: 261214},
			{input: []int{2, 3, 1}, expect: 6895259},
			{input: []int{3, 2, 1}, expect: 18},
			{input: []int{3, 1, 2}, expect: 362},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("Test%dExpect%d", i, test.expect), func(t *testing.T) {
				assert.Equal(t, test.expect, FindNumber(test.input, 30000000))
			})
		}
	})

}
