package day16

import (
	"errors"
	"fmt"
	"strings"

	bitutils "github.com/ralucas/advent-of-code/pkg/utils/bit"
)

type TypeID int

const (
	Operator TypeID = 1
	Literal  TypeID = 4
)

type LengthTypeID int

const (
	TotalBitLength LengthTypeID = 0
	NumSubPackets  LengthTypeID = 1
)

type PacketParser struct {
	raw    string
	rawArr []string
	*Packet
}

type Packet struct {
	version      int
	typeID       TypeID
	lengthTypeID LengthTypeID
}

func NewPacketParser(s string) *PacketParser {
	return &PacketParser{
		Packet: &Packet{},
		raw:    s,
	}
}

func (pp *PacketParser) parseHeader(bits []int8) error {
	if len(bits) < 7 {
		return errors.New(fmt.Sprintf("bad length of %d, should be 7", len(bits)))
	}

	pp.SetVersion(bits[:3])
	pp.SetTypeID(bits[3:6])
	pp.SetLengthTypeID(bits[7])

	return nil
}

func (pp *PacketParser) parseLiteral(bits []int8) {

	return
}

func (pp *PacketParser) Parse() (*Packet, error) {
	var bits []int8

	rawArr := strings.Split(pp.raw, "")

	for _, char := range rawArr {
		bin := bitutils.HexToBin[char]

		bits = append(bits, bin...)

		lb := len(bits)

		if lb >= 7 {
			if err := pp.parseHeader(bits[:7]); err != nil {
				return nil, err
			}
			bits = bits[7:]
			continue
		}

		if lb >= 15 && pp.TypeID() == Operator {
			switch pp.LengthTypeID() {
			case TotalBitLength:
				fmt.Print("tbl")
			case NumSubPackets:
				fmt.Print("nsp")
			}
		}
	}

	if pp.TypeID() == Literal {
		// handle literal
		pp.parseLiteral(bits)
	}

	return pp.Packet, nil
}

func (p *Packet) SetVersion(bitarr []int8) {
	p.version = bitutils.Btoi(bitarr)
}

func (p *Packet) Version() int {
	return p.version
}

func (p *Packet) SetTypeID(bitarr []int8) {
	val := bitutils.Btoi(bitarr)
	if val == 4 {
		p.typeID = Literal
		return
	}
	p.typeID = Operator
}

func (p *Packet) TypeID() TypeID {
	return p.typeID
}

func (p *Packet) SetLengthTypeID(bit int8) {
	p.lengthTypeID = LengthTypeID(bit)
}

func (p *Packet) LengthTypeID() LengthTypeID {
	return p.lengthTypeID
}
