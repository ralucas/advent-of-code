package day16

import bitutil "github.com/ralucas/advent-of-code/pkg/util/bit"

type TypeID int

const (
	Operator    TypeID = -1
	Sum         TypeID = 0
	Product     TypeID = 1
	Minimum     TypeID = 2
	Maximum     TypeID = 3
	Literal     TypeID = 4
	GreaterThan TypeID = 5
	LessThan    TypeID = 6
	Equal       TypeID = 7
)

type Packet struct {
	version      int
	typeID       TypeID
	lengthTypeID LengthTypeID
	literal      int // literal only
	numChildren  int // operator only
	children     []*Packet
	parent       *Packet
	value        int
}

func (p *Packet) AddChild() (child *Packet) {
	child = &Packet{parent: p}
	p.children = append(p.children, child)
	p.numChildren += 1

	return child
}

func (p *Packet) Children() []*Packet {
	return p.children
}

func (p *Packet) Parent() *Packet {
	return p.parent
}

func (p *Packet) SetValue(val int) {
	p.value = val
}

func (p *Packet) Value() int {
	return p.value
}

func (p *Packet) SetVersion(bitarr []int8) {
	p.version = bitutil.Btoi(bitarr)
}

func (p *Packet) Version() int {
	return p.version
}

func (p *Packet) TypeID() TypeID {
	return p.typeID
}

func (p *Packet) Literal() int {
	return p.literal
}

func (p *Packet) SetLiteral(val int) {
	p.literal = val
}

func (p *Packet) SetTypeID(bitarr []int8) {
	val := bitutil.Btoi(bitarr)

	typeIDs := []TypeID{
		Sum,
		Product,
		Minimum,
		Maximum,
		Literal,
		GreaterThan,
		LessThan,
		Equal,
	}

	if val < len(typeIDs) {
		p.typeID = typeIDs[val]
	} else {
		p.typeID = Operator
	}
}

func (p *Packet) SetLengthTypeID(bit int8) {
	p.lengthTypeID = LengthTypeID(bit)
}

func (p *Packet) LengthTypeID() LengthTypeID {
	return p.lengthTypeID
}

func (p *Packet) Evaluate() {
	switch p.typeID {
	case Sum:
		p.value = 0
		for _, child := range p.children {
			p.value += child.value
		}
	case Product:
		p.value = 1
		for _, child := range p.children {
			p.value *= child.value
		}
	case Minimum:
		p.value = p.children[0].value
		for _, child := range p.children {
			if child.value < p.value {
				p.value = child.value
			}
		}
	case Maximum:
		p.value = p.children[0].value
		for _, child := range p.children {
			if child.value > p.value {
				p.value = child.value
			}
		}
	case GreaterThan:
		if p.children[0].value > p.children[1].value {
			p.value = 1
		} else {
			p.value = 0
		}
	case LessThan:
		if p.children[0].value < p.children[1].value {
			p.value = 1
		} else {
			p.value = 0
		}
	case Equal:
		if p.children[0].value == p.children[1].value {
			p.value = 1
		} else {
			p.value = 0
		}
	}
}
