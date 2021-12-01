package day12

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var td Day

func init() {
	td.PrepareData("../../../test/testdata/2020/12/test_input.txt")
}

func TestPrepareData(t *testing.T) {
	t.Run("correct length", func(t *testing.T) {
		assert.Equal(t, 5, len(td.data))
	})

	t.Run("correct data", func(t *testing.T) {
		expects := []Navigation{
			{"F", 10},
			{"N", 3},
			{"F", 7},
			{"R", 90},
			{"F", 11},
		}
		for i := range td.data {
			assert.Equal(t, expects[i], td.data[i])
		}
	})
}

func TestShip_Turn(t *testing.T) {
	type test struct {
		input  Navigation
		expect Direction
	}

	sp := Position{E, 0, N, 0}
	ship := NewShip(sp)

	tests := []test{
		{input: Navigation{"R", 90}, expect: S},
		{input: Navigation{"L", 90}, expect: E},
		{input: Navigation{"L", 180}, expect: W},
		{input: Navigation{"R", 90}, expect: N},
	}

	for i := range tests {
		ship.Turn(tests[i].input.action, tests[i].input.value)
		assert.Equal(t, tests[i].expect, ship.GetPointing(), fmt.Sprintf("test %d input: %+v", i, tests[i].input))
	}
}

func TestPosition_MovePosition(t *testing.T) {
	type test struct {
		inputDir Direction
		inputVal int
		expect   Position
	}

	pos := Position{E, 0, N, 0}

	tests := []test{
		{inputDir: N, inputVal: 10, expect: Position{E, 0, N, 10}},
		{inputDir: W, inputVal: 10, expect: Position{W, 10, N, 10}},
		{inputDir: S, inputVal: 5, expect: Position{W, 10, N, 5}},
		{inputDir: E, inputVal: 10, expect: Position{W, 0, N, 5}},
	}

	for i, tt := range tests {
		pos.MovePosition(tt.inputDir, tt.inputVal)
		assert.Equal(t, tt.expect, pos, fmt.Sprintf("test %d, input dir %d, input val %d", i, tt.inputDir, tt.inputVal))
	}
}

func TestShip_Move(t *testing.T) {

	t.Run("correct positions", func(t *testing.T) {
		expects := []Position{
			{E, 10, N, 0},
			{E, 10, N, 3},
			{E, 17, N, 3},
			{E, 17, N, 3},
			{E, 17, S, 8},
		}
		sp := Position{E, 0, N, 0}
		ship := NewShip(sp)

		for i, nav := range td.data {
			ship.Move(nav)
			assert.Equal(t, expects[i], ship.GetPos(), fmt.Sprintf("test %d, nav %+v", i+1, nav))
		}
	})

	t.Run("correct pointing", func(t *testing.T) {
		facings := []Direction{E, E, E, S, S}
		sp := Position{E, 0, N, 0}
		ship := NewShip(sp)

		for i, nav := range td.data {
			ship.Move(nav)
			assert.Equal(t, facings[i], ship.pointing, fmt.Sprintf("test %d, nav %+v", i+1, nav))
		}
	})
}

func TestPosition_ManhattanDistance(t *testing.T) {
	sp := Position{E, 0, N, 0}
	ship := NewShip(sp)
	for _, nav := range td.data {
		ship.Move(nav)
	}

	lp := ship.GetPos()
	assert.Equal(t, 25, lp.ManhattanDistance())
}

func TestShip_TurnWaypoint(t *testing.T) {
	sp := Position{E, 0, N, 0}
	ship := NewShip(sp)
	ship.TurnWaypoint("R", 90)

	assert.Equal(t, Position{S, 10, E, 1}, ship.waypoint.pos)
}

func TestShip_MoveWithWaypoint(t *testing.T) {
	t.Run("ships position", func(t *testing.T) {
		expects := []Position{
			{E, 100, N, 10},
			{E, 100, N, 10},
			{E, 170, N, 38},
			{E, 170, N, 38},
			{E, 214, S, 72},
		}
		sp := Position{E, 0, N, 0}
		ship := NewShip(sp)
		for i, nav := range td.data {
			ship.MoveWithWaypoint(nav)
			assert.Equal(t, expects[i], ship.pos, fmt.Sprintf("test %d, nav %+v", i+1, nav))
		}
	})

	t.Run("waypoints position", func(t *testing.T) {
		expects := []Position{
			{E, 10, N, 1},
			{E, 10, N, 4},
			{E, 10, N, 4},
			{S, 10, E, 4},
			{S, 10, E, 4},
		}
		sp := Position{E, 0, N, 0}
		ship := NewShip(sp)
		for i, nav := range td.data {
			ship.MoveWithWaypoint(nav)
			assert.Equal(t, expects[i], ship.waypoint.pos, fmt.Sprintf("test %d, nav %+v", i+1, nav))
		}

	})
}
