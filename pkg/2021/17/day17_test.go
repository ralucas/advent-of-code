package day17_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day17 "github.com/ralucas/advent-of-code/pkg/2021/17"
)

var td day17.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/17/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	t.Skip("not started")
	result := td.Part1()
	expect := 45

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	t.Skip("not started")
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
