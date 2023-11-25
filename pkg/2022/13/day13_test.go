package day13_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	day13 "github.com/ralucas/advent-of-code/pkg/2022/13"
)

var td day13.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../../test/testdata/2022/13/test_input.txt")

	m.Run()
}

func TestNewListFromString(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		input := "[1,2,3,4,5,10]"
		vals := []int{1, 2, 3, 4, 5, 10}
		expect := day13.NewList(vals...)

		actual, err := day13.NewListFromString(input)
		require.Nil(t, err)

		assert.EqualValues(t, expect.Values(), actual.Values())
	})
}

func TestParse(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		input := "[1,2,3,4,5,10]"
		vals := []int{1, 2, 3, 4, 5, 10}

		actual, err := day13.Parse(input)
		require.Nil(t, err)

		assert.Equal(t, len(vals), len(actual))

		for i := range vals {
			assert.Equal(t, vals[i], actual[i].Value())
		}
	})

	t.Run("two array", func(t *testing.T) {
		input := "[[1,2],[3,4]]"
		vals := [][]int{{1, 2}, {3, 4}}

		actual, err := day13.Parse(input)
		require.Nil(t, err)

		assert.Equal(t, len(vals), len(actual))

		for i := range vals {
			for j := range vals[i] {
				assert.Equal(t, vals[i][j], actual[i].Children()[j].Value())
			}
		}
	})
}

func TestPart1(t *testing.T) {
	result := td.Part1()
	expect := 13

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
