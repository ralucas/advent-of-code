package day16

import (
	"log"
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	data string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFile(filepath)

	d.SetData(strings.TrimSpace(data))

	return
}

func (d *Day) SetData(s string) {
	d.data = s
}

func (d *Day) Part1() interface{} {
	pp := NewPacketParser(d.data)
	packets, err := pp.Parse()

	if err != nil {
		log.Fatalf("error in part 1 %+v\n", err)
	}

	result := 0

	for _, p := range packets {
		result += p.Version()
	}

	return result
}

func (d *Day) Part2() interface{} {
	return nil
}
