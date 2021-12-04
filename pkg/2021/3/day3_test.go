package day3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day3 "github.com/ralucas/advent-of-code/pkg/2021/3"
)

// TODO: Update this per implementation
//nolint
var td day3.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/3/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 198, td.Part1())
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 230, td.Part2())
}
