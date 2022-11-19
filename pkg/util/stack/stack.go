package stack

import "errors"

type Stack struct {
	values []string
	ptr    int
}

func NewStack(ss ...string) *Stack {
	return &Stack{
		values: ss,
	}
}

func (s *Stack) Empty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (string, error) {
	if s.Empty() {
		return "", errors.New("empty stack")
	}

	v := s.values[len(s.values)-1]

	s.values = s.values[:len(s.values)-1]

	return v, nil
}

func (s *Stack) Push(val string) {
	s.values = append(s.values, val)
}
