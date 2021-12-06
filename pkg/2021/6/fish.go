package day6

const (
	DaysToCreate        = 6
	NewFishDaysToCreate = 8
)

type Fish struct {
	startState   int
	prevState    int
	currentState int
}

func NewFish(start int) *Fish {
	return &Fish{start, start, start}
}

func (f *Fish) SetState(s int) {
	f.prevState = f.currentState
	if s < 0 {
		f.currentState = DaysToCreate
	} else {
		f.currentState = s
	}
}

func (f *Fish) CurrentState() int {
	return f.currentState
}

func (f *Fish) PrevState() int {
	return f.prevState
}

func (f *Fish) Created() bool {
	return f.prevState == 0
}
