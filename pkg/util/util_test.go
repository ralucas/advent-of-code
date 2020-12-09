package util

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	rf := ReadFile("../../test/testdata/3/test-input1.txt")
	assert.NotNil(t, rf)
}

func TestReadFileToArray(t *testing.T) {
	rfArray := ReadFileToArray("../../test/testdata/3/test-input1.txt", "\n")
	assert.NotNil(t, rfArray)
	assert.IsType(t, rfArray, []string{})
	assert.True(t, len(rfArray) > 1)
}

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

func TestQSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().Unix())
		testArr := rand.Perm(100)
		sorted := QSort(testArr)
		for x := 1; x < 100; x++ {
			assert.True(t, sorted[x-1] <= sorted[x])
		}
	}
}

func BenchmarkQSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().Unix())
		testArr := rand.Perm(100)
		QSort(testArr)
	}
}
