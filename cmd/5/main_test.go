package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInputExpects = []BoardingPass{
	BoardingPass{ 44, 5, 357 },
	BoardingPass{ 70, 7, 567 },
	BoardingPass{ 14, 7, 119 },
	BoardingPass{ 102, 4, 820 },
}

func TestPrepareData(t *testing.T) {
	data := prepareData("../../test/testdata/5/test_input.txt")

	assert.IsType(t, []string{}, data[0])
	assert.Equal(t, 4, len(data))
}

func TestGetRow(t *testing.T) {
	data := prepareData("../../test/testdata/5/test_input.txt")

	for i := 0; i < len(data); i++ {
		row := getRow(data[i][:7])
		assert.Equal(t, testInputExpects[i].row, row, fmt.Sprintf("input of %v\n", data[i][:7]))
	}
}

func TestGetCol(t *testing.T) {
	data := prepareData("../../test/testdata/5/test_input.txt")

	for i := 0; i < len(data); i++ {
		col := getCol(data[i][7:])
		assert.Equal(t, testInputExpects[i].col, col, fmt.Sprintf("input of %v\n", data[i][7:]))
	}
}

func TestToBoardingPass(t *testing.T) {
	data := prepareData("../../test/testdata/5/test_input.txt")

	for i := 0; i < len(data); i++ {
		bp := toBoardingPass(data[i])
		assert.Equal(t, testInputExpects[i], bp)
	}
}

func TestMaxId(t *testing.T) {
	data := prepareData("../../test/testdata/5/test_input.txt")

	maxId := getHighestSeatID(data)

	assert.Equal(t, 820, maxId)
}