package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	values map[T]bool
	size   int
}

func New[T comparable](ss ...T) *Set[T] {
	values := make(map[T]bool)

	for _, item := range ss {
		values[item] = true
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

func (s *Set[T]) Remove(val T) (removed bool) {
	if s.Empty() {
		return false
	}

	_, ok := s.values[val]

	return ok
}

func (s *Set[T]) Add(vals ...T) {
	count := 0

	for _, val := range vals {
		if _, ok := s.values[val]; !ok {
			s.values[val] = true
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
