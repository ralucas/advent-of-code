package day2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day2 "github.com/ralucas/advent-of-code/pkg/2021/2"
)

// TODO: Update this per implementation
//nolint
var td day2.Day

func TestMain(m *testing.M) {
	// TODO: Uncomment and utilize
	td.PrepareData("../../../test/testdata/2021/2/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 150, td.Part1())
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 900, td.Part2())
}
