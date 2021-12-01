package day12

type Ship struct {
	pos      Position
	start    Position
	pointing Direction
	waypoint *Waypoint
}

type Waypoint struct {
	pos   Position
	start Position
}

func NewShip(p Position) *Ship {
	return &Ship{
		pos:      p,
		start:    p,
		pointing: E,
		waypoint: NewWaypoint(Position{E, 10, N, 1}),
	}
}

func NewWaypoint(p Position) *Waypoint {
	return &Waypoint{
		pos:   p,
		start: p,
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

func (s *Ship) TurnWaypoint(turn string, deg int) {
	for deg > 0 {
		deg -= 90

		if turn == "L" {
			s.waypoint.pos.SetXDirection(s.waypoint.pos.xDirection - 1)
			s.waypoint.pos.SetYDirection(s.waypoint.pos.yDirection - 1)
		} else {
			s.waypoint.pos.SetXDirection(s.waypoint.pos.xDirection + 1)
			s.waypoint.pos.SetYDirection(s.waypoint.pos.yDirection + 1)
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

func (s *Ship) MoveWithWaypoint(n Navigation) {
	switch n.action {
	case "L":
		s.TurnWaypoint(n.action, n.value)
	case "R":
		s.TurnWaypoint(n.action, n.value)
	case "F":
		s.pos.MovePosition(s.waypoint.pos.xDirection, n.value*s.waypoint.pos.xUnits)
		s.pos.MovePosition(s.waypoint.pos.yDirection, n.value*s.waypoint.pos.yUnits)
	case "N":
		s.waypoint.pos.MovePosition(N, n.value)
	case "E":
		s.waypoint.pos.MovePosition(E, n.value)
	case "S":
		s.waypoint.pos.MovePosition(S, n.value)
	case "W":
		s.waypoint.pos.MovePosition(W, n.value)
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
