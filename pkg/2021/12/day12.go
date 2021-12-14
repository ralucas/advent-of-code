package day12

import (
	"log"
	"strings"
	"sync"
	"unicode"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	adjacencyList map[string][]string
	sync.RWMutex
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFileToArray(filepath, "\n")

	d.adjacencyList = make(map[string][]string)

	for _, line := range data {
		path := strings.Split(line, "-")
		from, to := path[0], path[1]

		if _, ok := d.adjacencyList[from]; !ok {
			d.adjacencyList[from] = make([]string, 0)
		}

		if _, ok := d.adjacencyList[to]; !ok {
			d.adjacencyList[to] = make([]string, 0)
		}

		d.adjacencyList[from] = append(d.adjacencyList[from], to)
		d.adjacencyList[to] = append(d.adjacencyList[to], from)
	}

	return
}

func (d *Day) Part1() interface{} {

	cb := func(p *Path) PathAllowanceVisitor {
		return func(vertex string) bool {
			if unicode.IsLower(rune(vertex[0])) {
				if _, ok := p.visited[vertex]; ok {
					return false
				}
			}
			return true
		}
	}

	paths := AllPaths("start", d.adjacencyList, WithAllowanceFunc(cb))

	return len(paths)
}

func (d *Day) Part2() interface{} {

	recorder := func(p *Path) PathVisitRecorder {
		return func(vertex string) {
			if unicode.IsLower(rune(vertex[0])) {
				if vertex != "start" && vertex != "end" {
					if p.visited[vertex] == 2 {
						d.Lock()
						p.info["small"] = "done"
						d.Unlock()
					}
					if p.visited[vertex] > 2 {
						log.Fatalf("vertex %s seen %d", vertex, p.visited[vertex])
					}
				}
			}
		}
	}

	gatekeeper := func(p *Path) PathAllowanceVisitor {
		return func(vertex string) bool {
			if unicode.IsLower(rune(vertex[0])) {
				if vertex == "start" || vertex == "end" {
					return false
				}

				if _, ok := p.info["small"]; ok {
					if _, ok := p.visited[vertex]; ok {
						return false
					}
				}
			}

			return true
		}
	}

	paths := AllPaths("start", d.adjacencyList, WithAllowanceFunc(gatekeeper), WithPathVisitRecorder(recorder))

	return len(paths)
}

func AllPaths(start string, adjList map[string][]string, opts ...PathOption) []*Path {
	nextNodes := adjList[start]

	var paths []*Path

	for _, node := range nextNodes {
		path := NewPath([]string{start}, opts...)
		paths = append(paths, dfs(node, adjList, path, opts...)...)
	}

	return paths
}

func dfs(node string, adjList map[string][]string, path *Path, opts ...PathOption) []*Path {
	var paths []*Path

	if node == "end" {
		path.End(node)
		paths = append(paths, path)
		return paths
	}

	if path.CanVisit(node) {
		path.Visit(node)
	} else {
		return nil
	}

	neighbors := adjList[node]

	for _, neighbor := range neighbors {
		newPath := NewPath(path.path, opts...)
		paths = append(paths, dfs(neighbor, adjList, newPath, opts...)...)
	}

	return paths
}
