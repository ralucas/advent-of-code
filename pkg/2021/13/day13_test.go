package day13_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day13 "github.com/ralucas/advent-of-code/pkg/2021/13"
)

var td day13.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/13/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 17

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := "#####\n#...#\n#...#\n#...#\n#####\n.....\n.....\n"

	assert.Equal(t, expect, result)
}
