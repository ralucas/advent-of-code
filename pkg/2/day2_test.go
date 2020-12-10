package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input  Password
	expect bool
}

func TestPrepareData(t *testing.T) {
	d := Day{}

	d.PrepareData("../../assets/2/input.txt")
	assert.IsType(t, d.data[1], Password{})
	assert.True(t, len(d.data) > 10)
}

func TestPasswordValidate(t *testing.T) {
	testPw := Password{1, 3, "a", "abcde"}
	valid := IsValid(testPw)
	assert.True(t, valid)
}

func TestPasswordValidByPos(t *testing.T) {
	testPws := []Test{
		{input: Password{1, 3, "a", "abcde"}, expect: true},
		{input: Password{1, 3, "b", "cdefg"}, expect: false},
		{input: Password{2, 9, "c", "ccccccccc"}, expect: false},
	}
	for _, tp := range testPws {
		valid := IsValidByPosition(tp.input)
		assert.Equal(t, valid, tp.expect)
	}
}
