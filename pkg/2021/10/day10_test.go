package day10_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day10 "github.com/ralucas/advent-of-code/pkg/2021/10"
)

var td day10.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/10/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 26397

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 288957

	assert.Equal(t, expect, result)
}
