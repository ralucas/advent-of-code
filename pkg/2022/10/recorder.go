package day10

import "fmt"

type Recorder struct {
	snapshots []CPUState
}

func NewRecorder() *Recorder {
	return &Recorder{
		snapshots: []CPUState{{0, 1, 0}},
	}
}

func (r *Recorder) Record(state CPUState) error {
	r.snapshots = append(r.snapshots, state)

	return nil
}

func (r *Recorder) RetrieveByCycle(cycle int) (CPUState, error) {
	if cycle >= len(r.snapshots) {
		return CPUState{}, fmt.Errorf("cycle %d is out of bounds", cycle)
	}

	return r.snapshots[cycle], nil
}

func (r *Recorder) Size() int {
	return len(r.snapshots)
}
