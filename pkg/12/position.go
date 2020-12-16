package day12

type Position struct {
	xDirection Direction
	xUnits     int
	yDirection Direction
	yUnits     int
}

func (p *Position) MovePosition(d Direction, units int) {
	switch d {
	case p.xDirection:
		p.xUnits += units
		break
	case p.yDirection:
		p.yUnits += units
		break
	case p.xDirection.Reverse():
		p.xUnits -= units
		if p.xUnits < 0 {
			p.xUnits = -p.xUnits
			p.xDirection = p.xDirection.Reverse()
		}
		break
	case p.yDirection.Reverse():
		p.yUnits -= units
		if p.yUnits < 0 {
			p.yUnits = -p.yUnits
			p.yDirection = p.yDirection.Reverse()
		}
		break
	}
}

func (p *Position) SetXDirection(d Direction) {
	if d < 0 {
		d = W
	}
	if d > 3 {
		d = N
	}

	p.xDirection = d
}

func (p *Position) SetYDirection(d Direction) {
	if d < 0 {
		d = W
	}
	if d > 3 {
		d = N
	}

	p.yDirection = d
}

func (p *Position) ManhattanDistance() int {
	return p.xUnits + p.yUnits
}
