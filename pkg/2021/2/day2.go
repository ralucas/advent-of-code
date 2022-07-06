package day2

import (
	"log"
	"strconv"
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	// TODO: Change this
	data []Command
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, "\n")

	for _, v := range data {
		spl := strings.Split(v, " ")

		moveUnits, err := strconv.Atoi(spl[1])
		if err != nil {
			log.Fatalf("could not convert string to int %+v", err)
		}

		c := Command{Movement(spl[0]), moveUnits}

		d.data = append(d.data, c)
	}

	return
}

func (d *Day) Part1() interface{} {
	sub := NewSub(0, 0, 0)
	for _, c := range d.data {
		sub.MoveNormal(c)
	}

	return sub.GetPos() * sub.GetDepth()
}

func (d *Day) Part2() interface{} {
	sub := NewSub(0, 0, 0)
	for _, c := range d.data {
		sub.MoveWithAim(c)
	}

	return sub.GetPos() * sub.GetDepth()
}
