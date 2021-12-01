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

var td Day

func init() {
	td.PrepareData("../../../test/testdata/2020/5/test_input.txt")
}

func TestPrepareData(t *testing.T) {

	assert.IsType(t, []string{}, td.data[0])
	assert.Equal(t, 4, len(td.data))
}

func TestGetRow(t *testing.T) {
	for i := 0; i < len(td.data); i++ {
		row := getRow(td.data[i][:7])
		assert.Equal(t, testInputExpects[i].Row, row, fmt.Sprintf("input of %v\n", td.data[i][:7]))
	}
}

func TestGetCol(t *testing.T) {
	for i := 0; i < len(td.data); i++ {
		col := getCol(td.data[i][7:])
		assert.Equal(t, testInputExpects[i].Col, col, fmt.Sprintf("input of %v\n", td.data[i][7:]))
	}
}

func TestToBoardingPass(t *testing.T) {
	for i := 0; i < len(td.data); i++ {
		bp := toBoardingPass(td.data[i])
		assert.Equal(t, testInputExpects[i], bp)
	}
}

func TestMaxId(t *testing.T) {
	maxId := GetHighestSeatID(td.data)

	assert.Equal(t, 820, maxId)
}

func TestNewPlane(t *testing.T) {
	p := NewPlane(td.data)

	t.Run("has correct row count", func(t *testing.T) {
		for i := 0; i < len(td.data); i++ {
			assert.Equal(t, 1, p.rows[testInputExpects[i].Row])
		}
	})

	t.Run("has correct plane row and col", func(t *testing.T) {
		for i := 0; i < len(td.data); i++ {
			assert.Equal(t, 1, p.seating[testInputExpects[i].Row][testInputExpects[i].Col])
		}
	})
}

func TestFindAvailableSeats(t *testing.T) {
	totalSeats := 128 * 8

	p := NewPlane(td.data)
	as := findAvailableSeats(p)
	assert.Equal(t, totalSeats-4, len(as))
}
