package day9_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day9 "github.com/ralucas/advent-of-code/pkg/2021/9"
)

var td day9.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/9/test_input.txt")

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
