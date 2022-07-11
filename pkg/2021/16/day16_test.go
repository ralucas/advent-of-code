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
	rootPacket, err := pp.Parse()

	require.Nil(t, err)

	assert.Equal(t, 2021, rootPacket.Value())
}

func TestOperatorPacketParserWithLength(t *testing.T) {
	pp := day16.NewPacketParser("38006F45291200")
	root, err := pp.Parse()

	require.Nil(t, err)

	assert.Equal(t, 2, len(root.Children()))
	assert.Equal(t, 10, root.Children()[0].Value())
	assert.Equal(t, 20, root.Children()[1].Value())
}

func TestOperatorPacketParserWithNumPackets(t *testing.T) {
	pp := day16.NewPacketParser("EE00D40C823060")
	root, err := pp.Parse()

	require.Nil(t, err)

	assert.Equal(t, 3, len(root.Children()))
	assert.Equal(t, 1, root.Children()[0].Value())
	assert.Equal(t, 2, root.Children()[1].Value())
	assert.Equal(t, 3, root.Children()[2].Value())
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
		{
			input:  "620080001611562C8802118E34",
			expect: 12,
		},
		{
			input:  "C0015000016115A2E0802F182340",
			expect: 23,
		},
		{
			input:  "A0016C880162017C3686B18A3D4780",
			expect: 31,
		},
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
	t.Parallel()

	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  "C200B40A82",
			expect: 3,
		},
		{
			input:  "04005AC33890",
			expect: 54,
		},
		{
			input:  "880086C3E88112",
			expect: 7,
		},
		{
			input:  "CE00C43D881120",
			expect: 9,
		},
		{
			input:  "D8005AC2A8F0",
			expect: 1,
		},
		{
			input:  "F600BC2D8F",
			expect: 0,
		},
		{
			input:  "9C005AC2F8F0",
			expect: 0,
		},
		{
			input:  "9C0141080250320F1802104A08",
			expect: 1,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			td.SetData(tc.input)
			result := td.Part2()

			assert.Equal(t, tc.expect, result)
		})
	}
}
