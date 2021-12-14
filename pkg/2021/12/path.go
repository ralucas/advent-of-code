package day12

import (
	"strings"
)

type Path struct {
	start         string
	end           string
	path          []string
	visited       map[string]int
	info          map[string]string
	gatekeeper    PathAllowanceVisitor
	visitRecorder PathVisitRecorder
}

type PathOption func(*Path)

type PathAllowanceVisitor func(string) bool
type PathAllowanceVisitorFactory func(*Path) PathAllowanceVisitor

type PathVisitRecorder func(string)
type PathVisitRecorderFactory func(*Path) PathVisitRecorder

func NewPath(nodes []string, opts ...PathOption) *Path {
	visited := make(map[string]int)

	for _, n := range nodes {
		visited[n] += 1
	}

	p := &Path{
		start:   nodes[0],
		path:    nodes,
		visited: visited,
		info:    make(map[string]string),
	}

	for _, opt := range opts {
		opt(p)
	}

	if p.visitRecorder != nil {
		for k, _ := range visited {
			p.visitRecorder(k)
		}
	}

	return p
}

func WithAllowanceFunc(fn PathAllowanceVisitorFactory) PathOption {
	return func(p *Path) {
		p.gatekeeper = fn(p)
	}
}

func WithPathVisitRecorder(fn PathVisitRecorderFactory) PathOption {
	return func(p *Path) {
		p.visitRecorder = fn(p)
	}
}

func (p *Path) CanVisit(vertex string) bool {
	if p.gatekeeper != nil && !p.gatekeeper(vertex) {
		return false
	}

	if p.Ended() {
		return false
	}

	return true
}

func (p *Path) Visit(vertex string) {
	if !p.CanVisit(vertex) {
		return
	}

	p.visited[vertex] += 1

	if p.visitRecorder != nil {
		p.visitRecorder(vertex)
	}

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

func (p *Path) String() string {
	var sb strings.Builder

	for _, n := range p.path {
		sb.WriteString(n)
		if n == "end" {
			break
		}
		sb.WriteString(" -> ")
	}

	return sb.String()
}
