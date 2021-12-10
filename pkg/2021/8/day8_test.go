package day8_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day8 "github.com/ralucas/advent-of-code/pkg/2021/8"
)

var td day8.Day
var td1 day8.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/8/test_input2.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 26

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		td1.PrepareData("../../../test/testdata/2021/8/test_input.txt")
		result := td1.Part2()
		expect := 5353

		assert.Equal(t, expect, result)
	})

	t.Run("Extensive", func(t *testing.T) {
		result := td.Part2()
		expect := 61229

		assert.Equal(t, expect, result)
	})
}
