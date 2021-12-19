package day15_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day15 "github.com/ralucas/advent-of-code/pkg/2021/15"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestPart1(t *testing.T) {
	t.Run("SmallSet", func(t *testing.T) {
		td := &day15.Day{}
		td.PrepareData("../../../test/testdata/2021/15/test_input.txt")
		result := td.Part1()
		expect := 40

		assert.Equal(t, expect, result)
	})

	t.Run("LargeSet", func(t *testing.T) {
		td2 := &day15.Day{}
		td2.PrepareData("../../../test/testdata/2021/15/test_input2.txt")
		result := td2.Part1()
		expect := 315

		assert.Equal(t, expect, result)
	})
}

func TestExpand(t *testing.T) {
	td := &day15.Day{}
	td.PrepareData("../../../test/testdata/2021/15/test_input.txt")
	grid := td.Builder().Expand(5).Build()

	td2 := &day15.Day{}
	td2.PrepareData("../../../test/testdata/2021/15/test_input2.txt")

	gridExpect := td2.Builder().Build()

	assert.ElementsMatch(t, gridExpect.Values(), grid.Values())
}

func TestPart2(t *testing.T) {
	td := &day15.Day{}
	td.PrepareData("../../../test/testdata/2021/15/test_input.txt")
	result := td.Part2()
	expect := 315

	assert.Equal(t, expect, result)
}
