package day7_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day7 "github.com/ralucas/advent-of-code/pkg/2022/7"
)

var td day7.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/7/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 95437
	
	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := 24933642

	assert.Equal(t, expect, result)
}
