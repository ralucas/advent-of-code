package day12

import (
	"log"
	"regexp"
	"strconv"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data []Navigation
}

type Navigation struct {
	action string
	value  int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	navs := make([]Navigation, len(data))
	reAlph := regexp.MustCompile(`[A-Za-z]+`)
	reNum := regexp.MustCompile(`[0-9]+`)
	for i, instruction := range data {
		action := reAlph.FindString(instruction)
		val, err := strconv.Atoi(reNum.FindString(instruction))
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		navs[i] = Navigation{action, val}
	}

	d.data = navs

	return
}

func (d *Day) Part1() interface{} {
	startingPos := Position{
		xDirection: E,
		xUnits:     0,
		yDirection: N,
		yUnits:     0,
	}
	ship := NewShip(startingPos)
	for _, nav := range d.data {
		ship.Move(nav)
	}
	finalPos := ship.GetPos()

	return finalPos.ManhattanDistance()
}

func (d *Day) Part2() interface{} {
	startingPos := Position{
		xDirection: E,
		xUnits:     0,
		yDirection: N,
		yUnits:     0,
	}
	ship := NewShip(startingPos)
	for _, nav := range d.data {
		ship.MoveWithWaypoint(nav)
	}
	finalPos := ship.GetPos()

	return finalPos.ManhattanDistance()
}
