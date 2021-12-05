package day5

import (
	"fmt"
	"log"
	"strings"

	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Day struct {
	lineBuilders []*LineBuilder
	maxX, maxY   int
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
		lb := NewLineBuilder(ptA, ptB)
		d.lineBuilders = append(d.lineBuilders, lb)

		if x := mathutils.Max(pA[0], pB[0]); x > d.maxX {
			d.maxX = x
		}

		if y := mathutils.Max(pA[1], pB[1]); y > d.maxY {
			d.maxY = y
		}
	}

	return
}

func printAllPoints(points [][]int) {
	for i := range points {
		for j := range points[i] {
			if points[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", points[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func (d *Day) calculateOverlaps(f func(l *LineBuilder) *Line) int {
	allPoints := make([][]int, d.maxY+1)
	for i := range allPoints {
		allPoints[i] = make([]int, d.maxY+1)
	}

	for _, lb := range d.lineBuilders {
		line := f(lb)
		for _, point := range line.Points() {
			allPoints[point.Y()][point.X()] += 1
		}
	}

	// printAllPoints(allPoints)

	counts := 0
	for i := range allPoints {
		for j := range allPoints[i] {
			if allPoints[i][j] > 1 {
				counts += 1
			}
		}
	}

	return counts
}

func (d *Day) Part1() interface{} {
	return d.calculateOverlaps(func(l *LineBuilder) *Line {
		return l.BuildWith().Horizontal().Vertical().BuildLine()
	})
}

func (d *Day) Part2() interface{} {
	return d.calculateOverlaps(func(l *LineBuilder) *Line {
		return l.BuildWith().Horizontal().Vertical().Diagonal().BuildLine()
	})
}
