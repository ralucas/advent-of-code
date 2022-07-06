package day20_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day20 "github.com/ralucas/advent-of-code/pkg/2021/20"
)

var td day20.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/20/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	t.Skip("not started")
	result := td.Part1()
	expect := true

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	t.Skip("not started")
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
