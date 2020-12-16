package day13

import (
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
