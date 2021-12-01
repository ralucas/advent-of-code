//go:build unit

package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var td Day

func init() {
	td.PrepareData("../../../test/testdata/2020/6/test_input.txt")
}

func TestPrepareData(t *testing.T) {
	assert.Equal(t, 5, len(td.data))
	assert.Equal(t, 3, len(td.data[1]))
	assert.Equal(t, 2, len(td.data[2]))
	assert.Equal(t, 4, len(td.data[3]))
}

func TestGroupCount(t *testing.T) {
	expects := []int{3, 3, 3, 1, 1}

	for i, d := range td.data {
		assert.Equal(t, expects[i], GroupCount(d))
	}
}

func TestSumGroupCounts(t *testing.T) {
	assert.Equal(t, 11, SumCounts(td.data, GroupCount))
}

func TestAllYesCount(t *testing.T) {
	expects := []int{3, 0, 1, 1, 1}

	for i, d := range td.data {
		assert.Equal(t, expects[i], AllYesCount(d))
	}
}

func TestSumAllYesCounts(t *testing.T) {
	assert.Equal(t, 6, SumCounts(td.data, AllYesCount))
}
