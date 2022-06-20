package day16_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	day16 "github.com/ralucas/advent-of-code/pkg/2021/16"
)

var td day16.Day

func TestMain(m *testing.M) {
	m.Run()
}

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  "8A004A801A8002F478",
			expect: 16,
		},
		// {
		// 	input:  "620080001611562C8802118E34",
		// 	expect: 12,
		// },
		// {
		// 	input:  "C0015000016115A2E0802F182340",
		// 	expect: 23,
		// },
		// {
		// 	input:  "A0016C880162017C3686B18A3D4780",
		// 	expect: 31,
		// },
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			td.SetData(tc.input)
			result := td.Part1()

			assert.Equal(t, tc.expect, result)
		})
	}
}

// func TestPart2(t *testing.T) {
// 	result := td.Part2()
// 	expect := true

// 	assert.Equal(t, expect, result)
// }
