package day13

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var td Day

func init() {
	td.PrepareData("../../test/testdata/13/test_input.txt")
}

func TestPrepareData(t *testing.T) {
	assert.Equal(t, 939, td.earliestTimestamp)
	assert.Equal(t, 8, len(td.busSchedule))
	assert.Equal(t, 5, len(td.buses))
}

func TestNearestNextBus(t *testing.T) {
	m, mb := NearestNextBus(td.earliestTimestamp, td.buses)

	assert.Equal(t, 5, m)
	assert.Equal(t, 59, mb)
}

func TestEarliestTimestampForScheduleSlow(t *testing.T) {
	m := EarliestTimestampForScheduleSlow(td.bsMap)

	assert.Equal(t, int64(1068781), m, fmt.Sprintf("expected %d, got %d", 1068781, m))
}

func TestEarliestTimestampForScheduleSlow2(t *testing.T) {
	type test struct {
		input  []map[int]int
		expect int64
	}

	testInputs := [][]string{
		{"17", "x", "13", "19"},
		{"67", "7", "59", "61"},
		{"67", "x", "7", "59", "61"},
		{"67", "7", "x", "59", "61"},
		{"1789", "37", "47", "1889"},
	}

	expects := []int64{3417, 754018, 779210, 1261476, 1202161486}

	tests := make([]test, len(testInputs))

	for i := range testInputs {
		tests[i].input = createScheduleMap(testInputs[i])
		tests[i].expect = expects[i]
	}

	for i, tt := range tests {
		name := fmt.Sprintf("test %d", i)
		t.Run(name, func(t *testing.T) {
			result := EarliestTimestampForScheduleSlow(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func TestIsPrime(t *testing.T) {
	type test struct {
		input  int
		expect bool
	}

	tests := []test{
		{input: 7, expect: true},
		{input: 5, expect: true},
		{input: 2, expect: true},
		{input: 3, expect: true},
		{input: 11, expect: true},
		{input: 41, expect: true},
		{input: 69, expect: false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, isPrime(tt.input))
	}
}

func TestPrimeFactors(t *testing.T) {
	type test struct {
		input  int
		expect []int
	}

	tests := []test{
		{input: 41, expect: []int{41}},
		{input: 12, expect: []int{2, 2, 3}},
		{input: 21, expect: []int{3, 7}},
		{input: 30, expect: []int{2, 3, 5}},
		{input: 45, expect: []int{3, 3, 5}},
		{input: 80, expect: []int{2, 2, 2, 2, 5}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, primeFactors(tt.input))
	}
}

func TestLeastCommonMultiples(t *testing.T) {
	type test struct {
		input  []int
		expect int
	}

	tests := []test{
		{input: []int{12, 80}, expect: 240},
		{input: []int{3, 9, 21}, expect: 63},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, leastCommonMultiples(tt.input))
	}
}
