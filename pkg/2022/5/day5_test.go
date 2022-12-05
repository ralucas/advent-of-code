package day5_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day5 "github.com/ralucas/advent-of-code/pkg/2022/5"
)

var td day5.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/5/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := "CMZ"

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := "MCD"

	assert.Equal(t, expect, result)
}
