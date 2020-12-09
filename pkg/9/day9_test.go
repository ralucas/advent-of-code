package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data []int

func init() {
	data = PrepareData("../../test/testdata/9/test_input.txt")
}

func TestPrepareData(t *testing.T) {
	assert.Equal(t, 20, len(data))
}

func TestFindFirstNonSum(t *testing.T) {
	f, _ := FindFirstNonSum(data, 5)
	assert.Equal(t, 127, f)
}

func TestContiguousSumSet(t *testing.T) {
	_, idx := FindFirstNonSum(data, 5)
	ss := ContiguousSumSet(data[:idx], 127)

	expects := []int{15, 25, 47, 40}

	for i, e := range expects {
		assert.Equal(t, e, ss[i])
	}
}

func TestMinMax(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}
	tests := []test{
		{input: []int{2, 1, 3, 4}, expect: []int{1, 4}},
		{input: []int{110, 10, 50, 55}, expect: []int{10, 110}},
		{input: []int{200, 2000, 20000, 10}, expect: []int{10, 20000}},
	}

	for _, tt := range tests {
		min, max := MinMax(tt.input)
		assert.Equal(t, tt.expect[0], min)
		assert.Equal(t, tt.expect[1], max)
	}
}
