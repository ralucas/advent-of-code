package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var td Day

func init() {
	td.PrepareData("../../test/testdata/8/test_input.txt")
}

func TestPrepareData(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		assert.NotNil(t, td.data)
	})

	t.Run("has the correct length", func(t *testing.T) {
		assert.Equal(t, 9, len(td.data))
	})

	t.Run("individual items are correct", func(t *testing.T) {
		testData := []Instruction{
			{op: "nop", sign: "+", val: 0},
			{op: "acc", sign: "+", val: 1},
			{op: "jmp", sign: "+", val: 4},
			{op: "acc", sign: "+", val: 3},
			{op: "jmp", sign: "-", val: 3},
			{op: "acc", sign: "-", val: 99},
			{op: "acc", sign: "+", val: 1},
			{op: "jmp", sign: "-", val: 4},
			{op: "acc", sign: "+", val: 6},
		}

		for i, test := range testData {
			assert.Equal(t, td.data[i].op, test.op)
			assert.Equal(t, td.data[i].sign, test.sign)
			assert.Equal(t, td.data[i].val, test.val)
		}
	})
}

func TestRunInstructions(t *testing.T) {

	t.Run("it returns correct last acc prior to looping", func(t *testing.T) {
		lastAcc, exitcode := RunInstructions(td.data)

		assert.Equal(t, 5, lastAcc)
		assert.Equal(t, -1, exitcode)
	})
}

func TestFixInstructions(t *testing.T) {
	acc := FixInstructions(td.data)

	assert.Equal(t, 8, acc)
}
