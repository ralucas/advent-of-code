package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	coordinates []Point
	directions  []Direction
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	file := fileutil.ReadFile(filepath)

	spl := strings.Split(file, "\n\n")

	rawCoords, rawDirs := spl[0], spl[1]

	coordLines := strings.Split(rawCoords, "\n")

	for _, line := range coordLines {
		s := arrayutil.MapToInt(strings.Split(line, ","))
		pt := Point{s[0], s[1]}
		d.coordinates = append(d.coordinates, pt)
	}

	dirLines := arrayutil.Filter(strings.Split(rawDirs, "\n"), func(s string) bool {
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
	grid.FoldPoints(d.directions[0])

	return len(grid.pointMap)
}

func (d *Day) Part2() interface{} {
	grid := NewGrid(d.coordinates).Build()

	for _, dir := range d.directions {
		grid.FoldPoints(dir)
	}

	vals := grid.PointsToGrid()
	ans := grid.String(vals)

	return fmt.Sprintf("\n%s", ans)
}
