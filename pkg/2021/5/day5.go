package day5

import (
	"log"
	"strings"

	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Day struct {
	lines      []Line
	maxX, maxY int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, "\n")

	for _, s := range data {
		points := strings.Split(s, " -> ")
		pA := arrayutils.MapToInt(strings.Split(points[0], ","))
		pB := arrayutils.MapToInt(strings.Split(points[1], ","))
		ptA := NewPoint(pA[0], pA[1])
		ptB := NewPoint(pB[0], pB[1])
		line := NewLine(ptA, ptB)
		d.lines = append(d.lines, line)

		if x := mathutils.Max(pA[0], pB[0]); x > d.maxX {
			d.maxX = x
		}

		if y := mathutils.Max(pA[1], pB[1]); y > d.maxY {
			d.maxY = y
		}
	}

	return
}

func (d *Day) Part1() interface{} {
	allPoints := make([][]int, d.maxX+1)
	for i := range allPoints {
		allPoints[i] = make([]int, d.maxY+1)
	}

	counts := 0
	for _, line := range d.lines {
		for _, point := range line.Points() {
			allPoints[point.X()][point.Y()] += 1
		}
	}

	for i := range allPoints {
		for j := range allPoints[i] {
			if allPoints[i][j] > 1 {
				counts += 1
			}
		}
	}

	return counts
}

func (d *Day) Part2() interface{} {
	return nil
}
