package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

var td, td1 Day

func init() {
	td.PrepareData("../../test/testdata/10/test_input.txt")
	td1.PrepareData("../../test/testdata/10/test_input1.txt")
}

func TestCountDiffs(t *testing.T) {

	t.Run("has correct counts for 2nd input", func(t *testing.T) {
		svi := utils.QSort(td1.data)
		upd := insertOutletAndDevice(svi)
		counts := CountDiffs(upd)

		assert.Equal(t, 22, counts[1])
		assert.Equal(t, 10, counts[3])
	})

}

func TestCountDistinctArrangements(t *testing.T) {

	t.Run("has the correct count on input 1", func(t *testing.T) {
		expected := [][]int{
			{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22},
			{0, 1, 4, 5, 6, 7, 10, 12, 15, 16, 19, 22},
			{0, 1, 4, 5, 7, 10, 11, 12, 15, 16, 19, 22},
			{0, 1, 4, 5, 7, 10, 12, 15, 16, 19, 22},
			{0, 1, 4, 6, 7, 10, 11, 12, 15, 16, 19, 22},
			{0, 1, 4, 6, 7, 10, 12, 15, 16, 19, 22},
			{0, 1, 4, 7, 10, 11, 12, 15, 16, 19, 22},
			{0, 1, 4, 7, 10, 12, 15, 16, 19, 22},
		}

		svi := utils.QSort(td.data)
		upd := insertOutletAndDevice(svi)
		_, tcount := BuildTree(upd)
		count := CountDistinctArrangements(upd)

		assert.Equal(t, len(expected), tcount, "Tree count failed")
		assert.Equal(t, len(expected), count, "Count failed")
	})

	t.Run("has the correct count on alternate input", func(t *testing.T) {
		input := []int{0, 1, 4, 5, 6, 7, 8, 10, 11, 12, 15, 16, 19, 22}

		_, tcount := BuildTree(input)
		count := CountDistinctArrangements(input)

		assert.Equal(t, tcount, count)
	})

	t.Run("has the correct count on input 2", func(t *testing.T) {
		svi := utils.QSort(td1.data)
		upd := insertOutletAndDevice(svi)
		_, tcount := BuildTree(upd)
		count := CountDistinctArrangements(upd)

		assert.Equal(t, 19208, tcount, "Tree count failed")
		assert.Equal(t, 19208, count, "Count failed")
	})

}
