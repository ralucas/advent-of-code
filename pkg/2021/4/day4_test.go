package day4_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day4 "github.com/ralucas/advent-of-code/pkg/2021/4"
)

var td day4.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/4/test_input.txt")

	m.Run()
}

func TestPrepareData(t *testing.T) {
	td.PrepareData("../../../test/testdata/2021/4/test_input.txt")

	assert.Equal(t, 27, len(td.Numbers()))

	for _, board := range td.Boards() {
		bvi := board.Values()
		for i := range bvi {
			assert.NotNil(t, bvi[i])
			for j := range bvi[i] {
				assert.NotNil(t, bvi[i][j])
			}
		}
	}
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 4512

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 1924

	assert.Equal(t, expect, result)
}
