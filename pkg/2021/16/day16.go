package day16

import (
	"log"
	"strings"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFile(filepath)

	d.SetData(strings.TrimSpace(data))
}

func (d *Day) SetData(s string) {
	d.data = s
}

func (d *Day) Part1() interface{} {
	pp := NewPacketParser(d.data)
	root, err := pp.Parse()

	if err != nil {
		log.Fatalf("error in part 1 %+v\n", err)
	}

	result := root.Version()

	children := root.children

	for len(children) > 0 {
		p := children[0]
		result += p.Version()
		children = append(children[1:], p.children...)
	}

	return result
}

func (d *Day) Part2() interface{} {
	pp := NewPacketParser(d.data)
	root, err := pp.Parse()

	if err != nil {
		log.Fatalf("error in part 2 %+v\n", err)
	}

	return root.Value()
}
