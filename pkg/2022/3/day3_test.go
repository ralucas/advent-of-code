package day3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day3 "github.com/ralucas/advent-of-code/pkg/2022/3"
)

var td day3.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/3/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 157

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 70

	assert.Equal(t, expect, result)
}
