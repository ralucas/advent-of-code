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
		_, count := DistinctArrangements(upd)
		//tree.Print()
		//count := CountPaths(tree.Root)

		assert.Equal(t, len(expected), count)
	})

	t.Run("has the correct count on input 2", func(t *testing.T) {
		svi := utils.QSort(td1.data)
		upd := insertOutletAndDevice(svi)
		_, count := DistinctArrangements(upd)
		//count := CountPaths(tree.Root, 0)
		//count := CountDistinctArrangements(upd)

		assert.Equal(t, 19208, count)
	})

}
