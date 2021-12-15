package day14_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day14 "github.com/ralucas/advent-of-code/pkg/2021/14"
)

var td day14.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/14/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 1588

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 2188189693529

	assert.Equal(t, expect, result)
}
