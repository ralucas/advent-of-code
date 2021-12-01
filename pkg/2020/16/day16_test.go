package day16

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint
var td Day
var td2 Day

func init() {
	td.PrepareData("../../../test/testdata/2020/16/test_input.txt")
	td2.PrepareData("../../../test/testdata/2020/16/test_input2.txt")
}

func TestPrepareData(t *testing.T) {
	t.Run("Lengths", func(t *testing.T) {
		assert.Equal(t, 3, len(td.fields))
		assert.Equal(t, 4, len(td.tickets))
		assert.Equal(t, 3, len(td.myticket.Vals))
	})

	t.Run("Tickets", func(t *testing.T) {
		for _, ticket := range td.tickets {
			assert.Equal(t, 3, len(ticket.Vals))
		}
	})
}

func TestBuildFieldFilter(t *testing.T) {
	t.Run("AllFields", func(t *testing.T) {
		ff := BuildFieldFilter(td.fields)
		assert.Equal(t, 51, len(ff))
		assert.Equal(t, 1, ff[len(ff)-1])
		assert.Equal(t, 0, ff[4])
		assert.Equal(t, 0, ff[0])
		assert.Equal(t, 0, ff[12])
	})

	t.Run("LimitFieldClass", func(t *testing.T) {
		ff := BuildFieldFilter(td.fields, "class")
		assert.Equal(t, 8, len(ff))
		assert.Equal(t, 1, ff[len(ff)-1])
		assert.Equal(t, 0, ff[4])
		assert.Equal(t, 0, ff[0])
	})
}

func TestSumInvalidTickets(t *testing.T) {
	ff := BuildFieldFilter(td.fields)

	sum := SumInvalidTickets(ff, td.tickets)

	assert.Equal(t, 71, sum)
}

func TestFindColumnsByFieldName(t *testing.T) {
	tests := []struct {
		input  string
		expect []int
	}{
		{input: "class", expect: []int{1, 2}},
		{input: "row", expect: []int{0, 1, 2}},
		{input: "seat", expect: []int{2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Field%s", test.input), func(t *testing.T) {
			c := FindColumnsByFieldName(td2, test.input)
			assert.Equal(t, test.expect, c)
		})
	}
}

func TestDiscoverColumns(t *testing.T) {
	colsmap := make(map[string][]int)
	for _, s := range []string{"class", "row", "seat"} {
		colsmap[s] = FindColumnsByFieldName(td2, s)
	}

	expect := map[string]int{
		"class": 1,
		"row":   0,
		"seat":  2,
	}

	assert.Equal(t, expect, DiscoverColumns(colsmap))
}
