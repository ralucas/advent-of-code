package math_test

import (
	"testing"

	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"

	"github.com/stretchr/testify/assert"
)

func TestSumg(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	floats := []float64{1.1, 2.2, 3.3, 4.4}

	t.Run("sum ints", func(t *testing.T) {
		assert.Equal(t, 10, mathutil.Sum(ints))
	})

	t.Run("sum floats", func(t *testing.T) {
		assert.Equal(t, float64(11.0), mathutil.Sum(floats))
	})
}