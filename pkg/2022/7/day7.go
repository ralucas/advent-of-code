package day7

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	terminalOutput []string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}

	d.terminalOutput = fileutil.ReadFileToArray(filepath, "\n")

	return
}

func (d *Day) Part1() interface{} {
	fs := NewFS()
	for _, s := range d.terminalOutput[1:] {
		fs.HandleOutput(s)
	}

	nodes := make([]*node, 0)

	for _, child := range fs.root.children {
		if child.IsDir() {
			nodes = append(nodes, child)
		}
	}

	total := 0

	for len(nodes) > 0 {
		n := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		if n.size <= 100000 {
			total += n.size
		}
		for _, child := range n.children {
			if child.IsDir() {
				nodes = append(nodes, child)
			}
		}
	}

	return total
}

func (d *Day) Part2() interface{} {
	totalSpace := 70000000
	minSpaceNeeded := 30000000

	fs := NewFS()
	for _, s := range d.terminalOutput[1:] {
		fs.HandleOutput(s)
	}

	availSpace := totalSpace - fs.root.size

	spaceNeeded := minSpaceNeeded - availSpace

	sizes := []int{fs.root.size}

	nodes := make([]*node, 0)

	for _, child := range fs.root.children {
		if child.IsDir() {
			nodes = append(nodes, child)
		}
	}

	for len(nodes) > 0 {
		n := nodes[len(nodes)-1]
		sizes = append(sizes, n.size)

		nodes = nodes[:len(nodes)-1]
		for _, child := range n.children {
			if child.IsDir() {
				nodes = append(nodes, child)
			}
		}
	}

	sortedSizes := sortutil.QSort(sizes)

	for _, size := range sortedSizes {
		if spaceNeeded-size <= 0 {
			return size
		}
	}

	return -1
}
