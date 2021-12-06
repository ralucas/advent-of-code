package day6_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	day6 "github.com/ralucas/advent-of-code/pkg/2021/6"
)

var td day6.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2021/6/test_input.txt")

	m.Run()
}

func TestPart1(t *testing.T) {
	t.Parallel()

	inputs := []int{18, 80}
	expects := []int{26, 5934}

	for i := range expects {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			td.SetDays(inputs[i])
			result := td.Part1()
			assert.Equal(t, expects[i], result)
		})
	}
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
