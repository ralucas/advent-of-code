package day12_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day12 "github.com/ralucas/advent-of-code/pkg/2022/12"
)

var td day12.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/12/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := true
	
	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
