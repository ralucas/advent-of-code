package day1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day1 "github.com/ralucas/advent-of-code/pkg/2022/1"
)

var td day1.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/1/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 24000

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 45000

	assert.Equal(t, expect, result)
}
