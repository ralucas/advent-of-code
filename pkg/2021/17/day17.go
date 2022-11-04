package day17

import (
	"log"
	"regexp"
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Day struct {
	xrange []int
	yrange []int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := strings.TrimSpace(fileutils.ReadFile(filepath))

	re := regexp.MustCompile(`=([-\d]+\.\.[-\d]+)`)

	matches := re.FindAllSubmatch([]byte(data), -1)

	d.xrange = arrayutils.MapToInt(strings.Split(string(matches[0][1]), ".."))
	d.yrange = arrayutils.MapToInt(strings.Split(string(matches[1][1]), ".."))
}

// The probe's x,y position starts at 0,0. Then, it will follow some trajectory by moving in steps. 
// On each step, these changes occur in the following order:
//
// The probe's x position increases by its x velocity.
// The probe's y position increases by its y velocity.
// Due to drag, the probe's x velocity changes by 1 toward the value 0; 
// 		that is, it decreases by 1 if it is greater than 0, increases by 1 
// 		if it is less than 0, or does not change if it is already 0.
// Due to gravity, the probe's y velocity decreases by 1.
func (d *Day) Part1() interface{} {
	xmax := mathutils.Max(d.xrange...)

	hx := xmax / 2

	xsteps := 0

	x := 0

	for hx > 0 {
		x += 1
		hx -= x
		xsteps += 1
	}

	ymax := mathutils.Max(d.yrange...)

	y := ymax

	start := 0

	height := y + (x-1) + (x-2) + (x-3)...(x-steps)

	return y
}

func (d *Day) Part2() interface{} {
	return nil
}
