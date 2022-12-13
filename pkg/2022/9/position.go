package day9

type Direction string

const (
	Right     Direction = "R"
	Left                = "L"
	Up                  = "U"
	Down                = "D"
	UpRight             = "UR"
	DownRight           = "DR"
	UpLeft              = "UL"
	DownLeft            = "DL"
)

type Move struct {
	direction Direction
	steps     int
}

type Position struct {
	x int
	y int
}

func NewStartPosition() *Position {
	return &Position{0, 0}
}

// func NewPositionLinkedList(n int) *Position {
// 	root := &Position{0, 0, nil}
// 	head := root
// 	for i := 0; i < n; i++ {
// 		head.next = &Position{0, 0, nil}
// 		head = head.next
// 	}

// 	return root
// }

func (p *Position) MakeStep(m Move) (updatedMove Move, done bool) {
	switch m.direction {
	case Right:
		p.x += 1
	case Left:
		p.x -= 1
	case Up:
		p.y += 1
	case Down:
		p.y -= 1
	case UpRight:
		p.y += 1
		p.x += 1
	case UpLeft:
		p.y += 1
		p.x -= 1
	case DownRight:
		p.y -= 1
		p.x += 1
	case DownLeft:
		p.y -= 1
		p.x -= 1
	}

	m.steps -= 1
	if m.steps == 0 {
		return m, true
	}
	return m, false
}

func (p *Position) Difference(other *Position) (x int, y int) {
	x = p.x - other.x
	y = p.y - other.y

	return x, y
}

func MovesFromDifference(x, y int) (moves []Move) {
	for x > 0 && y > 0 {
		m := Move{UpRight, 1}
		moves = append(moves, m)
		x -= 1
		y -= 1
	}

	for x > 0 && y < 0 {
		m := Move{DownRight, 1}
		moves = append(moves, m)
		x -= 1
		y += 1
	}

	for x < 0 && y > 0 {
		m := Move{UpLeft, 1}
		moves = append(moves, m)
		x += 1
		y -= 1
	}

	for x < 0 && y < 0 {
		m := Move{DownLeft, 1}
		moves = append(moves, m)
		x += 1
		y += 1
	}

	for x > 0 {
		m := Move{Right, 1}
		moves = append(moves, m)
		x -= 1
	}

	for x < 0 {
		m := Move{Left, 1}
		moves = append(moves, m)
		x += 1
	}

	for y > 0 {
		m := Move{Up, 1}
		moves = append(moves, m)
		y -= 1
	}

	for y < 0 {
		m := Move{Down, 1}
		moves = append(moves, m)
		y += 1
	}

	return moves
}
