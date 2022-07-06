package day16_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	day16 "github.com/ralucas/advent-of-code/pkg/2021/16"
)

var td day16.Day

func TestMain(m *testing.M) {
	m.Run()
}

func TestLiteralPacketParser(t *testing.T) {
	pp := day16.NewPacketParser("D2FE28")
	packets, err := pp.Parse()

	require.Nil(t, err)

	assert.Equal(t, 2021, packets[0].Literal())
}

func TestOperatorPacketParserWithLength(t *testing.T) {
	pp := day16.NewPacketParser("38006F45291200")
	packets, err := pp.Parse()

	require.Nil(t, err)

	assert.Equal(t, 3, len(packets))
	assert.Equal(t, 10, packets[1].Literal())
	assert.Equal(t, 20, packets[2].Literal())
}

func TestOperatorPacketParserWithNumPackets(t *testing.T) {
	pp := day16.NewPacketParser("EE00D40C823060")
	packets, err := pp.Parse()

	require.Nil(t, err)

	assert.Equal(t, 4, len(packets))
	assert.Equal(t, 1, packets[1].Literal())
	assert.Equal(t, 2, packets[2].Literal())
	assert.Equal(t, 3, packets[3].Literal())
}

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		expect int
	}{
		// {
		// 	input:  "8A004A801A8002F478",
		// 	expect: 16,
		// },
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

func TestPart2(t *testing.T) {
	t.Skip("not there")
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
