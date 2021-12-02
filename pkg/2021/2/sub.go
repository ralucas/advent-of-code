package day2

type Sub struct {
	pos   int
	depth int
	aim   int
}

func NewSub(pos, depth, aim int) *Sub {
	return &Sub{
		pos, depth, aim,
	}
}

func (s *Sub) GetPos() int {
	return s.pos
}

func (s *Sub) GetDepth() int {
	return s.depth
}

func (s *Sub) GetAim() int {
	return s.aim
}

func (s *Sub) MoveNormal(c Command) {
	switch c.movement {
	case Forward:
		s.pos += c.units
	case Back:
		s.pos -= c.units
	case Down:
		s.depth += c.units
	case Up:
		s.depth -= c.units
	}
}

func (s *Sub) MoveWithAim(c Command) {
	switch c.movement {
	case Forward:
		s.pos += c.units
		s.depth += (c.units * s.aim)
	case Back:
		s.pos -= c.units
	case Down:
		s.aim += c.units
	case Up:
		s.aim -= c.units
	}
}
