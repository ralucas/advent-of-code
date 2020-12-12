package day11

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var td Day
var outputData = make([][][]string, 5)

func TestMain(m *testing.M) {
	td.PrepareData("../../test/testdata/11/test_input.txt")

	for i := 0; i < 5; i++ {
		d := i + 1
		tdo := Day{}
		fp := fmt.Sprintf("../../test/testdata/11/test_output%d.txt", d)
		tdo.PrepareData(fp)
		data := make([][]string, len(tdo.data))
		copy(data, tdo.data)
		outputData[i] = data
	}

	m.Run()
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
			filled := Fill(inputData)
			for i := range outputData[n] {
				for j := range outputData[n][i] {
					assert.Equal(t, outputData[n][i][j], filled[i][j], fmt.Sprintf("Failed 1st round, i: %d, j: %d", i, j))
				}
			}
		})
	}
}

func TestFillAllRounds(t *testing.T) {
	filled := td.data

	for n := range outputData {
		filled = Fill(filled)
		for i := range outputData[n] {
			for j := range outputData[n][i] {
				assert.Equal(t, outputData[n][i][j], filled[i][j], fmt.Sprintf("Failed n: %d, i: %d, j: %d", n, i, j))
			}
		}
	}
}

func TestIsEqual(t *testing.T) {
	assert.True(t, IsEqual(outputData[0], Fill(td.data)))
	assert.False(t, IsEqual(outputData[0], td.data))
}

func TestCountSeats(t *testing.T) {
	occCount, _ := CountSeats(outputData[len(outputData)-1])
	assert.Equal(t, 37, occCount)
}
