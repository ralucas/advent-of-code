package array_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ralucas/advent-of-code/pkg/util/array"
)

func TestMapTo2D(t *testing.T) {
	testArr := []string{"1:A", "2:B", "3:C"}
	arr2D := array.MapTo2D(testArr, ":")
	for _, vi := range arr2D {
		assert.Equal(t, 2, len(vi))
	}
}

func TestMapToInt(t *testing.T) {
	testArr := []string{"1", "2", "3"}
	mappedInts := array.MapToInt(testArr)
	for i, v := range mappedInts {
		assert.Equal(t, i+1, v)
	}
}

func TestFilter(t *testing.T) {
	testArr := []string{"1", "2", "3"}
	filtered := array.Filter(testArr, func(s string) bool { return s != "2" })
	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, "1", filtered[0])
	assert.Equal(t, "3", filtered[1])
}

func TestFilterInt(t *testing.T) {
	testArr := []int{1, 2, 3}
	filtered := array.FilterInt(testArr, func(v int) bool { return v != 2 })
	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, 1, filtered[0])
	assert.Equal(t, 3, filtered[1])
}

func TestFindIntIndexes(t *testing.T) {
	testArr := []int{1, 2, 3}
	foundIdxs := array.FindIntIndexes(testArr, func(v int) bool { return v == 2 })
	assert.Equal(t, 1, len(foundIdxs))
	assert.Equal(t, 1, foundIdxs[0])
}

func TestEvery(t *testing.T) {
	testArr := []int{2, 4, 6, 8}

	t.Run("Every returns true", func(t *testing.T) {
		every := array.Every(testArr, func(v int, i int) bool { return v%2 == 0 })
		assert.True(t, every)
	})

	t.Run("Every returns false", func(t *testing.T) {
		every := array.Every(testArr, func(v int, i int) bool { return v-2 == 0 })
		assert.False(t, every)
	})
}

func TestIndex(t *testing.T) {
	t.Parallel()
	inputArr := []string{"a", "b", "c", "d", "e"}
	tests := []struct {
		input  string
		expect int
	}{
		{input: "c", expect: 2},
		{input: "f", expect: -1},
		{input: "a", expect: 0},
		{input: "e", expect: 4},
		{input: "hasdf", expect: -1},
		{input: "xdasc", expect: -1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			assert.Equal(t, test.expect, array.Index(inputArr, test.input))
		})
	}
}

func TestIndexInt(t *testing.T) {
	t.Parallel()
	inputArr := []int{1, 2, 3, 4, 5}
	tests := []struct {
		input  int
		expect int
	}{
		{input: 3, expect: 2},
		{input: 20, expect: -1},
		{input: 1, expect: 0},
		{input: 5, expect: 4},
		{input: 1234, expect: -1},
		{input: 100, expect: -1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			assert.Equal(t, test.expect, array.IndexInt(inputArr, test.input))
		})
	}
}

func TestIndexesInt(t *testing.T) {
	t.Parallel()
	inputArr := []int{1, 1, 1, 2, 3, 4, 5, 2, 3}
	tests := []struct {
		input  int
		expect []int
	}{
		{input: 3, expect: []int{4, 8}},
		{input: 20, expect: nil},
		{input: 1, expect: []int{0, 1, 2}},
		{input: 5, expect: []int{6}},
		{input: 1234, expect: nil},
		{input: 100, expect: nil},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			assert.Equal(t, test.expect, array.IndexesInt(inputArr, test.input))
		})
	}
}

func TestTranspose(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	expect := [][]int{
		{1, 6},
		{2, 7},
		{3, 8},
		{4, 9},
		{5, 10},
	}

	result := array.Transpose(input)

	assert.Equal(t, expect, result)
}

func TestDiagonal(t *testing.T) {
	t.Run("ErrorsOnNonSquare", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{4, 5, 6, 7},
		}

		_, err := array.Diagonals(input)
		assert.Error(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}

		expect := [][]int{
			{1, 7, 13, 19, 25},
			{21, 17, 13, 9, 5},
		}

		result, err := array.Diagonals(input)
		require.NoError(t, err)

		assert.Equal(t, expect, result)
	})
}
