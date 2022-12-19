package day11_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day11 "github.com/ralucas/advent-of-code/pkg/2022/11"
)

var td day11.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/11/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 10605

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 2713310158

	assert.Equal(t, expect, result)
}
