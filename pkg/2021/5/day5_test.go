package day5_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	day5 "github.com/ralucas/advent-of-code/pkg/2021/5"
)

func TestLine(t *testing.T) {
	input := [][]day5.Point{
		{day5.NewPoint(0, 0), day5.NewPoint(0, 2)},
		{day5.NewPoint(0, 0), day5.NewPoint(2, 0)},
		{day5.NewPoint(5, 8), day5.NewPoint(2, 8)},
	}

	expect := [][]day5.Point{
		{day5.NewPoint(0, 0), day5.NewPoint(0, 1), day5.NewPoint(0, 2)},
		{day5.NewPoint(0, 0), day5.NewPoint(1, 0), day5.NewPoint(2, 0)},
		{day5.NewPoint(5, 8), day5.NewPoint(4, 8), day5.NewPoint(3, 8), day5.NewPoint(2, 8)},
	}

	for i := range input {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			line := day5.NewLine(input[i][0], input[i][1])
			for j, p := range line.Points() {
				assert.Equal(t, expect[i][j], p)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	var td day5.Day

	td.PrepareData("../../../test/testdata/2021/5/test_input.txt")
	result := td.Part1()
	expect := 5

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	var td day5.Day

	td.PrepareData("../../../test/testdata/2021/5/test_input.txt")
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
