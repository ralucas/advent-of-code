package utils

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		n := 1000
		testArr := rand.Perm(n)
		sorted := QSort(testArr)
		for x := 1; x < n; x++ {
			assert.True(t, sorted[x-1] <= sorted[x])
		}
	}
}

func BenchmarkQSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		n := 1000
		testArr := rand.Perm(n)
		QSort(testArr)
	}
}

func TestMergeSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		n := 1000
		testArr := rand.Perm(n)
		sorted := MergeSort(testArr)
		for x := 1; x < n; x++ {
			assert.True(t, sorted[x-1] <= sorted[x])
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		n := 1000
		testArr := rand.Perm(n)
		MergeSort(testArr)
	}
}
