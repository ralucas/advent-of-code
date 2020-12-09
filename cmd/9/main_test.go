package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data []int

func init() {
	data = prepareData("../../test/testdata/9/test_input.txt")
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
