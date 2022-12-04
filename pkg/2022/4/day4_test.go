package day4_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day4 "github.com/ralucas/advent-of-code/pkg/2022/4"
)

var td day4.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/4/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 2

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 4

	assert.Equal(t, expect, result)
}
