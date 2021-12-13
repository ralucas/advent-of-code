package day12

import "unicode"

type Path struct {
	start   string
	end     string
	path    []string
	visited map[string]int
}

func NewPath(start, end string) *Path {
	visited := make(map[string]int)

	visited[start] = 1

	return &Path{
		start:   start,
		path:    []string{start},
		visited: visited,
	}
}

func (p *Path) Visit(vertex string) {
	if unicode.IsLower(rune(vertex[0])) {
		if _, ok := p.visited[vertex]; ok {
			return
		}
	}

	p.visited[vertex] += 1
	p.path = append(p.path, vertex)
}

func (p *Path) End(node string) {
	p.Visit(node)
	p.end = node
}

func (p *Path) Ended() bool {
	return p.end != ""
}

func (p *Path) Length() int {
	return len(p.path)
}

func (p *Path) Last() string {
	return p.path[len(p.path)-1]
}
