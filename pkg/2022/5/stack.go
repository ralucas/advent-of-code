package day5

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	values []string
	ptr    int
	size   int
}

func NewStack(ss ...string) *Stack {
	return &Stack{
		values: ss,
		size:   len(ss),
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Pop() (string, error) {
	if s.Empty() {
		return "", errors.New("empty stack")
	}

	v := s.values[s.size-1]

	s.values = s.values[:s.size-1]

	s.size -= 1

	return v, nil
}

func (s *Stack) PopN(n int) ([]string, error) {
	if s.Empty() {
		return nil, errors.New("empty stack")
	}

	if s.size < n {
		return nil, errors.New(fmt.Sprintf("stack size [%d] less than requested [%d]", s.size, n))
	}

	popIdx := s.size - n

	vals := s.values[popIdx:]

	s.values = s.values[:popIdx]

	s.size -= n

	return vals, nil
}

func (s *Stack) Push(vals ...string) {
	s.values = append(s.values, vals...)

	s.size += len(vals)
}

func (s *Stack) Peek() string {
	if s.Empty() {
		return ""
	}

	return s.values[s.size-1]
}

func (s *Stack) Print() string {
	var sb strings.Builder

	sb.WriteString("[")
	for _, v := range s.values {
		ws := fmt.Sprintf(" %s", v)
		sb.WriteString(ws)
	}
	sb.WriteString(" ]")

	return sb.String()
}