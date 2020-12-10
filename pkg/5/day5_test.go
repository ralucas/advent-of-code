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

var testDay Day

func init() {
	testDay.PrepareData("../../test/testdata/5/test_input.txt")
}

func TestPrepareData(t *testing.T) {

	assert.IsType(t, []string{}, testDay.data[0])
	assert.Equal(t, 4, len(testDay.data))
}

func TestGetRow(t *testing.T) {
	for i := 0; i < len(testDay.data); i++ {
		row := getRow(testDay.data[i][:7])
		assert.Equal(t, testInputExpects[i].Row, row, fmt.Sprintf("input of %v\n", testDay.data[i][:7]))
	}
}

func TestGetCol(t *testing.T) {
	for i := 0; i < len(testDay.data); i++ {
		col := getCol(testDay.data[i][7:])
		assert.Equal(t, testInputExpects[i].Col, col, fmt.Sprintf("input of %v\n", testDay.data[i][7:]))
	}
}

func TestToBoardingPass(t *testing.T) {
	for i := 0; i < len(testDay.data); i++ {
		bp := toBoardingPass(testDay.data[i])
		assert.Equal(t, testInputExpects[i], bp)
	}
}

func TestMaxId(t *testing.T) {
	maxId := GetHighestSeatID(testDay.data)

	assert.Equal(t, 820, maxId)
}

func TestNewPlane(t *testing.T) {
	p := NewPlane(testDay.data)

	t.Run("has correct row count", func(t *testing.T) {
		for i := 0; i < len(testDay.data); i++ {
			assert.Equal(t, 1, p.rows[testInputExpects[i].Row])
		}
	})

	t.Run("has correct plane row and col", func(t *testing.T) {
		for i := 0; i < len(testDay.data); i++ {
			assert.Equal(t, 1, p.seating[testInputExpects[i].Row][testInputExpects[i].Col])
		}
	})
}

func TestFindAvailableSeats(t *testing.T) {
	totalSeats := 128 * 8

	p := NewPlane(testDay.data)
	as := findAvailableSeats(p)
	assert.Equal(t, totalSeats-4, len(as))
}
