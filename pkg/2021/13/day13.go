package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
	mathutils "github.com/ralucas/advent-of-code/pkg/utils/math"
)

type Day struct {
	coordinates []*Point
	directions  []Direction
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	file := fileutils.ReadFile(filepath)

	spl := strings.Split(file, "\n\n")

	rawCoords, rawDirs := spl[0], spl[1]

	coordLines := strings.Split(rawCoords, "\n")

	for _, line := range coordLines {
		s := arrayutils.MapToInt(strings.Split(line, ","))
		pt := NewPoint(s[0], s[1])
		d.coordinates = append(d.coordinates, pt)
	}

	dirLines := arrayutils.Filter(strings.Split(rawDirs, "\n"), func(s string) bool {
		return s != ""
	})

	for _, line := range dirLines {
		s := strings.Split(line, "=")
		val, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}

		dn := Direction{
			text:  line,
			axis:  string(s[0][len(s[0])-1]),
			value: val,
		}

		d.directions = append(d.directions, dn)
	}

	return
}

func (d *Day) Part1() interface{} {
	grid := NewGrid(d.coordinates).Build()
	fold := grid.Fold(d.directions[0])

	count := 0

	for i := range fold {
		count += mathutils.Sum(fold[i])
	}

	return count
}

func (d *Day) Part2() interface{} {
	grid := NewGrid(d.coordinates).Build()

	for _, dir := range d.directions {
		grid.Fold(dir)
	}

	ans := grid.String(grid.foldValues)

	return fmt.Sprintf("\n%s", ans)
}
