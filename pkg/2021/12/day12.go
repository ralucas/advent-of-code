package day12

import (
	"log"
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	adjacencyList map[string][]string
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, "\n")

	for _, line := range data {
		path := strings.Split(line, "-")
		from, to := path[0], path[1]
		d.adjacencyList[from] = append(d.adjacencyList[from], to)
		d.adjacencyList[to] = append(d.adjacencyList[to], from)
	}

	return
}

func (d *Day) Part1() interface{} {
	return nil
}

func (d *Day) Part2() interface{} {
	return nil
}


// DFS takes a start, end points and an adjacency list and outputs
// a list of distinct paths.
func DepthFirstSearch(start, end string, adjList map[string][]string) [][]string {
	stack := NewStack(start)

	for !stack.Empty() {
		node, _ := stack.Pop()

		nextNodes := adjList[node]

		for _, path := range paths {

		}

		for _, nn := range nextNodes {

		}

	}
}