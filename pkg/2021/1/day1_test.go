//go:build unit

package day1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day1 "github.com/ralucas/advent-of-code/pkg/2021/1"
)

var td day1.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/1/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	c := td.Part1()

	assert.Equal(t, 7, c)
}

func TestPart2(t *testing.T) {
	c := td.Part2()

	assert.Equal(t, 5, c)
}
