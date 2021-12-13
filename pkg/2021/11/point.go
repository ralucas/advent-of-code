package day11

type Point struct {
	row        int
	col        int
	value      Value
	flashCount int
	didFlash   bool
	observers []*Point
}

type Value struct {
	initial  int
	previous int
	current  int
}

func NewPoint(row, col, value int) *Point {
	return &Point{
		row:        row,
		col:        col,
		value:      NewValue(value),
		flashCount: 0,
	}
}

func NewValue(val int) Value {
	return Value{val, val, val}
}

// 1. The energy level of each octopus increases by 1.
// 2. Any octopus with an energy level greater than 9 flashes.
// 		This increases the energy level of all adjacent octopuses by 1,
// 		including octopuses that are diagonally adjacent. If this
//		causes an octopus to have an energy level greater than 9,
//		it also flashes. This process continues as long as new octopuses
//		keep having their energy level increased beyond 9.
//	  (An octopus can only flash at most once per step.)
// 3. Finally, any octopus that flashed during this step has its energy
// 		level set to 0, as it used all of its energy to flash.
func (p *Point) Add(value int) {
	if p.didFlash {
		return
	}

	p.value.previous = p.value.current

	p.value.current += value

	if p.value.current > 9 {
		p.Flash()
	}

	p.value.current = p.value.current % 10

}

func (p *Point) RegisterObserver(o *Point) {
	// first need to make sure it's
	// not already registered
	// TODO: there's a better way
	// could use a map, but is there
	// an alternative here
	for _, co := range p.observers {
		if co == o {
			return 
		}
	}

	p.observers = append(p.observers, o)
}

func (p *Point) CurrentValue() int {
	return p.value.current
}

func (p *Point) Flash() {
	p.flashCount += 1
	p.didFlash = true
	p.notify()
}

func (p *Point) DidFlash() bool {
	return p.didFlash
}

func (p *Point) FlashCount() int {
	return p.flashCount
}

func (p *Point) Reset() {
	p.didFlash = false
}

func (p *Point) notify() {
	for _, o := range p.observers {
		o.Add(1)
	}
}