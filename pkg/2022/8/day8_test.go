package day8_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day8 "github.com/ralucas/advent-of-code/pkg/2022/8"
)

var td day8.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/8/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 21

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 8

	assert.Equal(t, expect, result)
}
