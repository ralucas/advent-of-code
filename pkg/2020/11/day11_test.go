//go:build unit

package day11

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var td, td1, td2 Day
var outputData = make([][][]string, 5)
var outputData2 = make([][][]string, 6)

func init() {
	td.PrepareData("../../../test/testdata/2020/11/part1/test_input.txt")

	for i := 0; i < 5; i++ {
		d := i + 1
		tdo := Day{}
		fp := fmt.Sprintf("../../../test/testdata/2020/11/part1/test_output%d.txt", d)
		tdo.PrepareData(fp)
		data := make([][]string, len(tdo.data))
		copy(data, tdo.data)
		outputData[i] = data
	}

	td1.PrepareData("../../../test/testdata/2020/11/part2/test_input.txt")
	td2.PrepareData("../../../test/testdata/2020/11/part2/test_input_no_replace.txt")

	for i := 0; i < 6; i++ {
		d := i + 1
		tdo := Day{}
		fp := fmt.Sprintf("../../../test/testdata/2020/11/part2/test_output%d.txt", d)
		tdo.PrepareData(fp)
		data := make([][]string, len(tdo.data))
		copy(data, tdo.data)
		outputData2[i] = data
	}
}

func TestPrepareData(t *testing.T) {
	assert.Equal(t, 10, len(td.data))
	assert.Equal(t, 10, len(td.data[0]))
}

func TestGetSeatState(t *testing.T) {

	t.Run("success with specified data", func(t *testing.T) {
		testInput := [][]string{
			{"#", ".", "#", "#", ".", "L", "#", ".", "#", "#"},
			{"#", "L", "#", "#", "#", "L", "L", ".", "L", "#"},
			{"L", ".", "#", ".", "#", ".", ".", "#", ".", "."},
		}

		expects := []string{"#", "L", "L", "L", "#", "L", "L", ".", "L", "#"}

		i := 1
		for j := range testInput[i] {
			ss := GetSeatState(i, j, testInput)
			assert.Equal(t, expects[j], ss)
		}
	})

	t.Run("success with first row of input", func(t *testing.T) {
		for j := range outputData[0][0] {
			ss := GetSeatState(0, j, td.data)
			assert.Equal(t, outputData[0][0][j], ss)
		}
	})

}

func TestFill(t *testing.T) {

	for n := 0; n < 5; n++ {
		name := fmt.Sprintf("run %d round", n+1)
		var inputData [][]string
		if n == 0 {
			inputData = td.data
		} else {
			inputData = outputData[n-1]
		}
		t.Run(name, func(t *testing.T) {
			filled := Fill(inputData, GetSeatState)
			for i := range outputData[n] {
				for j := range outputData[n][i] {
					assert.Equal(t, outputData[n][i][j], filled[i][j], fmt.Sprintf("Failed %d round, i: %d, j: %d", n+1, i, j))
				}
			}
		})
	}
}

func TestFillAllRounds(t *testing.T) {
	filled := td.data

	for n := range outputData {
		filled = Fill(filled, GetSeatState)
		for i := range outputData[n] {
			for j := range outputData[n][i] {
				assert.Equal(t, outputData[n][i][j], filled[i][j], fmt.Sprintf("Failed n: %d, i: %d, j: %d", n+1, i, j))
			}
		}
	}
}

func TestIsEqual(t *testing.T) {
	assert.True(t, IsEqual(outputData[0], Fill(td.data, GetSeatState)))
	assert.False(t, IsEqual(outputData[0], td.data))
}

func TestCountSeats(t *testing.T) {
	occCount, _ := CountSeats(outputData[len(outputData)-1])
	assert.Equal(t, 37, occCount)
}

func TestGetVisualSeatState(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		vis := GetVisualSeatState(4, 3, td2.data)

		assert.Equal(t, td2.data[4][3], vis)
	})

	t.Run("2", func(t *testing.T) {
		vis := GetVisualSeatState(0, 6, outputData2[0])

		assert.Equal(t, "L", vis)
	})
}

func TestFillVisualAllRounds(t *testing.T) {
	filled := td1.data

	for n := range outputData {
		filled = Fill(filled, GetVisualSeatState)
		for i := range outputData2[n] {
			for j := range outputData2[n][i] {
				assert.Equal(t, outputData2[n][i][j], filled[i][j], fmt.Sprintf("Failed n: %d, i: %d, j: %d", n+1, i, j))
			}
		}
	}
}

func TestFillVisualState(t *testing.T) {
	for n := 0; n < 6; n++ {
		name := fmt.Sprintf("run %d round", n+1)
		var inputData [][]string
		if n == 0 {
			inputData = td1.data
		} else {
			inputData = outputData2[n-1]
		}
		t.Run(name, func(t *testing.T) {
			filled := Fill(inputData, GetVisualSeatState)
			for i := range outputData2[n] {
				for j := range outputData2[n][i] {
					assert.Equal(t, outputData2[n][i][j], filled[i][j], fmt.Sprintf("Failed %d round, i: %d, j: %d", n+1, i, j))
				}
			}
		})
	}

}
