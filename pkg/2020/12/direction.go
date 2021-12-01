package day12

import "fmt"

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (d Direction) String() string {
	return []string{"N", "E", "S", "W"}[d]
}

func (d Direction) Reverse() Direction {
	switch d {
	case N:
		return S
	case E:
		return W
	case S:
		return N
	case W:
		return E
	}

	return d
}

func NewDirectionFromString(dir string) (Direction, error) {
	sDirs := []string{"N", "E", "S", "W"}
	for i := range sDirs {
		if sDirs[i] == dir {
			return Direction(i), nil
		}
	}

	return -1, fmt.Errorf("No such direction: %s", dir)
}
