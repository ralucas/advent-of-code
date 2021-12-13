package day11

import "errors"

type Stack struct {
	values []*Point
	ptr    int
}

func NewStack(pts ...*Point) *Stack {
	return &Stack{
		values: pts,
	}
}

func (s *Stack) Empty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (*Point, error) {
	if s.Empty() {
		return nil, errors.New("empty stack")
	}

	v := s.values[len(s.values)-1]

	s.values = s.values[:len(s.values)-1]

	return v, nil
}

func (s *Stack) Push(p *Point) {
	s.values = append(s.values, p)
}
