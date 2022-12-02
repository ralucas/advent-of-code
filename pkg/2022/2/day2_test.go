package day2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day2 "github.com/ralucas/advent-of-code/pkg/2022/2"
)

var td day2.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/2/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 15
	
	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 12

	assert.Equal(t, expect, result)
}
