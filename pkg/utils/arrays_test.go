package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapToInt(t *testing.T) {
	testArr := []string{"1", "2", "3"}
	mappedInts := MapToInt(testArr)
	for i, v := range mappedInts {
		assert.Equal(t, i+1, v)
	}
}

func TestFilter(t *testing.T) {
	testArr := []string{"1", "2", "3"}
	filtered := Filter(testArr, func(s string) bool { return s != "2" })
	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, "1", filtered[0])
	assert.Equal(t, "3", filtered[1])
}

func TestFilterInt(t *testing.T) {
	testArr := []int{1, 2, 3}
	filtered := FilterInt(testArr, func(v int) bool { return v != 2 })
	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, 1, filtered[0])
	assert.Equal(t, 3, filtered[1])
}

func TestFindIntIndexes(t *testing.T) {
	testArr := []int{1, 2, 3}
	foundIdxs := FindIntIndexes(testArr, func(v int) bool { return v == 2 })
	assert.Equal(t, 1, len(foundIdxs))
	assert.Equal(t, 1, foundIdxs[0])
}

func TestEvery(t *testing.T) {
	testArr := []int{2, 4, 6, 8}

	t.Run("Every returns true", func(t *testing.T) {
		every := Every(testArr, func(v int) bool { return v%2 == 0 })
		assert.True(t, every)
	})

	t.Run("Every returns false", func(t *testing.T) {
		every := Every(testArr, func(v int) bool { return v-2 == 0 })
		assert.False(t, every)
	})
}

func TestIndex(t *testing.T) {
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
			assert.Equal(t, test.expect, Index(inputArr, test.input))
		})
	}

}
