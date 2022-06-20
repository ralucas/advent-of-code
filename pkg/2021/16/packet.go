package day16

import (
	"strings"

	"github.com/pkg/errors"

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
	raw     string
	rawArr  []string
	packets []*Packet
}

type Packet struct {
	version      int
	typeID       TypeID
	lengthTypeID LengthTypeID
	literal      int
}

func NewPacketParser(s string) *PacketParser {
	return &PacketParser{
		packets: []*Packet{},
		raw:     s,
	}
}

func (pp *PacketParser) parseHeader(packet *Packet, bits []int8) (*Packet, error) {
	if len(bits) < 7 {
		return nil, errors.Errorf("bad length of %d, should be 7", len(bits))
	}

	packet.SetVersion(bits[:3])
	packet.SetTypeID(bits[3:6])
	packet.SetLengthTypeID(bits[6])

	return packet, nil
}

func (pp *PacketParser) parseLiteral(packet *Packet, bits []int8) error {
	packet.literal = bitutils.Btoi(bits)

	return nil
}

func (pp *PacketParser) parseOperator(packet *Packet, bits []int8) error {
	if packet.LengthTypeID() == TotalBitLength {
		totalLength := bitutils.Btoi(bits[:15])
		bits = bits[15:]
		pkt := new(Packet)

		pp.packets = append(pp.packets, pkt)

		if totalLength > len(bits) {
			totalLength = len(bits)
		}

		return pp.parsePacket(pkt, bits[:totalLength])
	}

	if packet.LengthTypeID() == NumSubPackets {
		n := bitutils.Btoi(bits[:11])
		bits = bits[11:]
		bits = trimLeadingZeros(bits)
		pktlen := len(bits) / n

		for i := 0; i < n; i++ {
			pkt := new(Packet)

			pp.packets = append(pp.packets, pkt)

			start, end := i*pktlen, (i+1)*pktlen

			err := pp.parsePacket(pkt, bits[start:end])
			if err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New("failed to parse operator")
}

func (pp *PacketParser) parsePacket(packet *Packet, bits []int8) error {
	packet, err := pp.parseHeader(packet, bits[:7])
	if err != nil {
		return err
	}

	bits = bits[7:]

	if packet.TypeID() == Literal {
		return pp.parseLiteral(packet, bits)
	}

	if packet.TypeID() == Operator {
		return pp.parseOperator(packet, bits)
	}

	return errors.New("failed to parse packet")
}

func (pp *PacketParser) Parse() ([]*Packet, error) {
	var bits []int8
	var err error

	rawArr := strings.Split(pp.raw, "")

	packet := new(Packet)

	pp.packets = append(pp.packets, packet)

	for _, char := range rawArr {
		bin := bitutils.HexToBin[char]

		bits = append(bits, bin...)
	}

	bits = trimTrailingZeros(bits)

	err = pp.parsePacket(packet, bits)
	if err != nil {
		return nil, err
	}

	return pp.packets, nil
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

func trimTrailingZeros(bits []int8) []int8 {
	count := 0

	for i := len(bits) - 1; i >= 0; i-- {
		if bits[i] == 0 {
			count++
		} else {
			break
		}
	}

	return bits[:len(bits)-count]
}

func trimLeadingZeros(bits []int8) []int8 {
	count := 0

	for i := 0; i < len(bits); i++ {
		if bits[i] == 0 {
			count++
		} else {
			break
		}
	}

	return bits[count:]
}
