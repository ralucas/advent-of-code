package day9_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	day9 "github.com/ralucas/advent-of-code/pkg/2022/9"
)

var td day9.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/9/test_input.txt")

	m.Run()
}

func TestAllVisited(t *testing.T) {
	td.Part1()
	av := td.AllVisited()

	expected := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{4, 2},
		{4, 3},
		{3, 4},
		{2, 4},
		{3, 3},
		{4, 3},
		{3, 2},
		{2, 2},
		{1, 2},
	}

	assert.EqualValues(t, expected, av)
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 13

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		expect    int
	}{
		{
			name:      "basic",
			inputFile: "test_input.txt",
			expect:    1,
		},
		{
			name:      "complex",
			inputFile: "test_input2.txt",
			expect:    36,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testDay := &day9.Day{}
			testDay.PrepareData(fmt.Sprintf("../../../test/testdata/2022/9/%s", tc.inputFile))
			result := testDay.Part2()
			assert.Equal(t, tc.expect, result)
		})
	}
}
