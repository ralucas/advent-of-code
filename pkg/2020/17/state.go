package day17

import "log"

type State int

const (
	Active State = iota
	Inactive
)

func NewState(state string) State {
	switch state {
	case "#":
		return Active
	case ".":
		return Inactive
	default:
		log.Fatalf("bad state: [%s]", state)
	}

	return -1
}

func (s State) String() string {
	return []string{"#", "."}[s]
}
