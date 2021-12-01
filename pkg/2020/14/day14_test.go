package day14

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var td Day

func init() {
	td.PrepareData("../../../test/testdata/2020/14/test_input.txt")
}

func TestPrepareData(t *testing.T) {
	t.Run("CorrectLength", func(t *testing.T) {
		assert.Equal(t, 1, len(td.data))
	})

	t.Run("CorrectData", func(t *testing.T) {
		p := ProgramSet{
			mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			mem: []Memory{
				{loc: 8, val: 11, valbin: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1}},
				{loc: 7, val: 101, valbin: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1}},
				{loc: 8, val: 0, valbin: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			},
		}

		assert.Equal(t, p, td.data[0])
	})
}

func TestItob(t *testing.T) {
	type test struct {
		input  int
		expect []int8
	}

	tests := []test{
		{input: 1, expect: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{input: 2, expect: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}},
		{input: 3, expect: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}},
		{input: 8, expect: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}},
		{input: 10, expect: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0}},
	}

	for i, v := range tests {
		v := v
		name := fmt.Sprintf("Test%d__input%d", i, v.input)
		t.Run(name, func(t *testing.T) {
			b := itob(v.input)
			assert.Equal(t, v.expect, b)
		})
	}
}

func TestBtoi(t *testing.T) {
	type test struct {
		input  []int8
		expect int
	}

	tests := []test{
		{expect: 1, input: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{expect: 2, input: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}},
		{expect: 3, input: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}},
		{expect: 8, input: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}},
		{expect: 10, input: []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0}},
	}

	for i, v := range tests {
		v := v
		name := fmt.Sprintf("Test%d", i)
		t.Run(name, func(t *testing.T) {
			b := btoi(v.input)
			assert.Equal(t, v.expect, b)
		})
	}
}

func TestApplyMask(t *testing.T) {
	expects := []int{73, 101, 64}

	for i, expect := range expects {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			am := ApplyMask(td.data[0].mem[i].valbin, td.data[0].mask)
			assert.Equal(t, expect, btoi(am))
		})
	}
}

func TestPart1(t *testing.T) {
	o := td.Part1()
	assert.Equal(t, int64(165), o)
}

func TestCountFloatingBits(t *testing.T) {
	mask := "100X0X11X0X"

	assert.Equal(t, 4, CountFloatingBits(mask))
}

func TestBitPermutations(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		p := BitPermutations(1)
		var expect = [][]int8{{0}, {1}}

		assert.Equal(t, expect, p)
	})

	t.Run("2", func(t *testing.T) {
		p := BitPermutations(2)
		var expect = [][]int8{{0, 0}, {1, 0}, {0, 1}, {1, 1}}

		assert.Equal(t, expect, p)
	})

	t.Run("3", func(t *testing.T) {
		p := BitPermutations(3)
		var expect = [][]int8{
			{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {1, 1, 0},
			{0, 0, 1}, {1, 0, 1}, {0, 1, 1}, {1, 1, 1},
		}

		assert.Equal(t, expect, p)
	})
}

func TestPart2(t *testing.T) {
	td2 := Day{}
	td2.PrepareData("../../../test/testdata/2020/14/test_input2.txt")

	out := td2.Part2()

	assert.Equal(t, int64(208), out)
}
