package day6

import (
	"fmt"
	"strings"
)

type State struct {
	fish []*Fish
}

func NewState(starts []int) *State {
	fishes := make([]*Fish, len(starts))
	for i, start := range starts {
		fishes[i] = NewFish(start)
	}

	return &State{fishes}
}

func (s *State) Day() {
	for _, f := range s.fish {
		f.SetState(f.CurrentState() - 1)
		if f.Created() {
			s.fish = append(s.fish, NewFish(NewFishDaysToCreate))
		}
	}
}

func (s *State) FishCount() int {
	return len(s.fish)
}

func (s *State) Print() string {
	var sb strings.Builder
	for _, f := range s.fish {
		sb.WriteString(fmt.Sprintf("%d", f.CurrentState()))
		sb.WriteString(",")
	}

	return sb.String()
}
