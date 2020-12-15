package day12

type Ship struct {
	pos      Position
	start    Position
	pointing Direction
}

func NewShip(p Position) *Ship {
	return &Ship{
		pos:      p,
		start:    p,
		pointing: E,
	}
}

func (s *Ship) Turn(turn string, deg int) {
	for deg > 0 {
		deg -= 90

		if turn == "L" {
			s.SetPointing(s.pointing - 1)
		} else {
			s.SetPointing(s.pointing + 1)
		}

		if s.pointing < 0 {
			s.SetPointing(W)
		}

		if s.pointing > 3 {
			s.SetPointing(N)
		}
	}
}

func (s *Ship) Move(n Navigation) {
	switch n.action {
	case "L":
		s.Turn(n.action, n.value)
	case "R":
		s.Turn(n.action, n.value)
	case "F":
		s.pos.MovePosition(s.pointing, n.value)
	case "N":
		s.pos.MovePosition(N, n.value)
	case "E":
		s.pos.MovePosition(E, n.value)
	case "S":
		s.pos.MovePosition(S, n.value)
	case "W":
		s.pos.MovePosition(W, n.value)
	}
}

func (s *Ship) GetPos() Position {
	return s.pos
}

func (s *Ship) SetPos(p Position) {
	s.pos = p
}

func (s *Ship) GetPointing() Direction {
	return s.pointing
}

func (s *Ship) SetPointing(d Direction) {
	s.pointing = d
}
