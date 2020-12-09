package utils

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
