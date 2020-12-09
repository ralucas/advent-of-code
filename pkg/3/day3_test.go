package day3

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/ralucas/advent-of-code/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestPrepareData(t *testing.T) {
	d := PrepareData("../../assets/3/input.txt")
	assert.Equal(t, 323, len(d))
	assert.IsType(t, d[0], []string{})
}

func TestSetGetPos(t *testing.T) {
	type test struct {
		input  int
		expect int
	}

	startState := SledState{
		Start: 0,
		Pos:   0,
		End:   4,
		Right: 3,
		Down:  1,
	}
	tests := []test{
		{input: 5, expect: 5},
		{input: 2, expect: 2},
		{input: 5, expect: 5},
		{input: 1, expect: 1},
		{input: 6, expect: 6},
		{input: 9, expect: 9},
	}
	for _, tt := range tests {
		startState.SetPos(tt.input)
		assert.Equal(t, tt.expect, startState.GetPos())
	}
}

func TestIsEqualToPosition(t *testing.T) {
	type test struct {
		input  []string
		expect bool
	}

	startState := SledState{
		Start: 0,
		Pos:   0,
		End:   4,
		Right: 3,
		Down:  1,
	}
	tests := []test{
		{input: []string{".", "#", "#", "#", "#"}, expect: false},
		{input: []string{".", "#", "#", "#", "#"}, expect: true},
		{input: []string{".", "#", "#", "#", "#"}, expect: true},
		{input: []string{".", "#", "#", "#", "#"}, expect: true},
		{input: []string{".", "#", "#", "#", "#"}, expect: true},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expect, startState.IsEqualToPosition(tt.input, "#"))
		startState.SetPos(1 + startState.GetPos())
	}
}

func TestGetNextPosition(t *testing.T) {
	startState := SledState{
		Start: 0,
		Pos:   0,
		End:   len([]string{".", "#", "#", "#", "#", ".", ".", "#"}) - 1,
		Right: 3,
		Down:  1,
	}

	tests := []int{3, 6, 1, 4, 7, 2}
	for _, tt := range tests {
		np := startState.NextPosition()
		assert.Equal(t, tt, np)
		startState.SetPos(np)
	}
}

func TestIsTree(t *testing.T) {
	type test struct {
		input  []string
		expect bool
	}

	var tests []test

	data := util.ReadFile("../../test/testdata/3/test-input.txt")
	lData := strings.Split(data, "\n")

	lineLength := 0

	for _, line := range lData {
		l := strings.Split(line, ",")
		b, _ := strconv.ParseBool(l[1])
		ll := strings.Split(l[0], "")
		lineLength = len(ll)
		tests = append(tests, test{
			input:  ll,
			expect: b,
		})
	}

	startState := SledState{
		Start: 0,
		Pos:   0,
		End:   lineLength - 1,
		Right: 3,
		Down:  1,
	}
	for i, tt := range tests {
		b := startState.IsEqualToPosition(tt.input, "#")
		assert.Equal(t, tt.expect, b, fmt.Sprintf("Failed on line %d, incorrect pos of %d", i, startState.GetPos()))
		startState.SetPos(startState.NextPosition())
	}
}

func TestMultiplesIsTree(t *testing.T) {
	type test struct {
		input  [][]string
		slope  []int
		expect int
	}

	tests := []test{
		{input: PrepareData("../../test/testdata/3/test-input2.txt"), slope: []int{1, 1}, expect: 2},
		{input: PrepareData("../../test/testdata/3/test-input2.txt"), slope: []int{3, 1}, expect: 7},
		{input: PrepareData("../../test/testdata/3/test-input2.txt"), slope: []int{5, 1}, expect: 3},
		{input: PrepareData("../../test/testdata/3/test-input2.txt"), slope: []int{7, 1}, expect: 4},
		{input: PrepareData("../../test/testdata/3/test-input2.txt"), slope: []int{1, 2}, expect: 2},
	}

	total := 1
	for i, tt := range tests {
		startState := SledState{
			Start: 0,
			Pos:   0,
			End:   len(tt.input[0]) - 1,
			Right: tt.slope[0],
			Down:  tt.slope[1],
		}

		treeCount := 0
		for j := 0; j < len(tt.input); j += startState.Down {
			if startState.IsEqualToPosition(tt.input[j], "#") {
				treeCount += 1
			}
			startState.SetPos(startState.NextPosition())
		}
		total *= treeCount
		assert.Equal(t, tt.expect, treeCount,
			fmt.Sprintf("Failed on input %d, with right %d and down %d", i, startState.Right, startState.Down))
	}
	assert.Equal(t, 336, total)
}
