package day5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputExpects = []BoardingPass{
	BoardingPass{44, 5, 357},
	BoardingPass{70, 7, 567},
	BoardingPass{14, 7, 119},
	BoardingPass{102, 4, 820},
}

func TestPrepareData(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	assert.IsType(t, []string{}, data[0])
	assert.Equal(t, 4, len(data))
}

func TestGetRow(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	for i := 0; i < len(data); i++ {
		row := getRow(data[i][:7])
		assert.Equal(t, testInputExpects[i].Row, row, fmt.Sprintf("input of %v\n", data[i][:7]))
	}
}

func TestGetCol(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	for i := 0; i < len(data); i++ {
		col := getCol(data[i][7:])
		assert.Equal(t, testInputExpects[i].Col, col, fmt.Sprintf("input of %v\n", data[i][7:]))
	}
}

func TestToBoardingPass(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	for i := 0; i < len(data); i++ {
		bp := toBoardingPass(data[i])
		assert.Equal(t, testInputExpects[i], bp)
	}
}

func TestMaxId(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	maxId := GetHighestSeatID(data)

	assert.Equal(t, 820, maxId)
}

func TestNewPlane(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	p := NewPlane(data)

	t.Run("has correct row count", func(t *testing.T) {
		for i := 0; i < len(data); i++ {
			assert.Equal(t, 1, p.rows[testInputExpects[i].Row])
		}
	})

	t.Run("has correct plane row and col", func(t *testing.T) {
		for i := 0; i < len(data); i++ {
			assert.Equal(t, 1, p.seating[testInputExpects[i].Row][testInputExpects[i].Col])
		}
	})
}

func TestFindAvailableSeats(t *testing.T) {
	data := PrepareData("../../test/testdata/5/test_input.txt")

	totalSeats := 128 * 8

	p := NewPlane(data)
	as := findAvailableSeats(p)
	assert.Equal(t, totalSeats-4, len(as))
}
