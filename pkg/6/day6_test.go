package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareData(t *testing.T) {
	data := PrepareData("../../test/testdata/6/test_input.txt")

	assert.Equal(t, 5, len(data))
	assert.Equal(t, 3, len(data[1]))
	assert.Equal(t, 2, len(data[2]))
	assert.Equal(t, 4, len(data[3]))
}

func TestGroupCount(t *testing.T) {
	data := PrepareData("../../test/testdata/6/test_input.txt")

	expects := []int{3, 3, 3, 1, 1}

	for i, d := range data {
		assert.Equal(t, expects[i], GroupCount(d))
	}
}

func TestSumGroupCounts(t *testing.T) {
	data := PrepareData("../../test/testdata/6/test_input.txt")

	assert.Equal(t, 11, SumCounts(data, GroupCount))
}

func TestAllYesCount(t *testing.T) {
	data := PrepareData("../../test/testdata/6/test_input.txt")

	expects := []int{3, 0, 1, 1, 1}

	for i, d := range data {
		assert.Equal(t, expects[i], AllYesCount(d))
	}
}

func TestSumAllYesCounts(t *testing.T) {
	data := PrepareData("../../test/testdata/6/test_input.txt")

	assert.Equal(t, 6, SumCounts(data, AllYesCount))
}
