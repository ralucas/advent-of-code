package day16

import bitutils "github.com/ralucas/advent-of-code/pkg/utils/bit"

type TypeID int

const (
	Operator TypeID = 1
	Literal  TypeID = 4
)

type Packet struct {
	version      int
	typeID       TypeID
	lengthTypeID LengthTypeID
	literal      int // literal only
	numChildren  int // operator only
	children     []*Packet
}

func (p *Packet) SetVersion(bitarr []int8) {
	p.version = bitutils.Btoi(bitarr)
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

func (p *Packet) SetTypeID(bitarr []int8) {
	val := bitutils.Btoi(bitarr)

	if val == int(Literal) {
		p.typeID = Literal
		return
	}

	p.typeID = Operator
}

func (p *Packet) SetLengthTypeID(bit int8) {
	p.lengthTypeID = LengthTypeID(bit)
}

func (p *Packet) LengthTypeID() LengthTypeID {
	return p.lengthTypeID
}
