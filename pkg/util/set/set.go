package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	values map[T]struct{}
	size   int
}

func New[T comparable](ss ...T) *Set[T] {
	values := make(map[T]struct{})

	for _, item := range ss {
		values[item] = struct{}{}
	}

	return &Set[T]{
		values: values,
		size:   len(values),
	}
}

// Empty if the set is empty.
func (s *Set[T]) Empty() bool {
	return s.size == 0
}

// Size of the set.
func (s *Set[T]) Size() int {
	return s.size
}

func (s *Set[T]) Remove(val T) {
	if s.Empty() {
		return
	}

	_, ok := s.values[val]
	if ok {
		delete(s.values, val)
		s.size -= 1
	}

	return
}

func (s *Set[T]) Add(vals ...T) {
	count := 0

	for _, val := range vals {
		if _, ok := s.values[val]; !ok {
			s.values[val] = struct{}{}
			count += 1
		}
	}

	s.size += count
}

func (s *Set[T]) Has(element T) (exists bool) {
	_, ok := s.values[element]

	return ok
}

func (s *Set[T]) Print() string {
	var sb strings.Builder

	sb.WriteString("[")
	for k := range s.values {
		ws := fmt.Sprintf(" %+v", k)
		sb.WriteString(ws)
	}
	sb.WriteString(" ]")

	return sb.String()
}
