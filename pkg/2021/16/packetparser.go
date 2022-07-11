package day16

import (
	"strings"

	"github.com/pkg/errors"

	bitutils "github.com/ralucas/advent-of-code/pkg/utils/bit"
)

type LengthTypeID int

const (
	TotalBitLength LengthTypeID = 0
	NumSubPackets  LengthTypeID = 1
)

type PacketParser struct {
	raw          string
	root         *Packet
	lengthTypeID LengthTypeID
}

func NewPacketParser(s string) *PacketParser {
	return &PacketParser{
		root: &Packet{},
		raw:  s,
	}
}

func (pp *PacketParser) parseHeader(packet *Packet, bits []int8) (*Packet, error) {
	if len(bits) < 7 {
		return nil, errors.Errorf("bad length of %d, should be 7", len(bits))
	}

	packet.SetVersion(bits[:3])
	packet.SetTypeID(bits[3:6])

	if packet.TypeID() != Literal {
		packet.SetLengthTypeID(bits[6])
	}

	return packet, nil
}

// Literal value packets encode a single binary number.
// To do this, the binary number is padded with leading
// zeroes until its length is a multiple of four bits,
// and then it is broken into groups of four bits. Each
// group is prefixed by a 1 bit except the last group,
// which is prefixed by a 0 bit. These groups of five
// bits immediately follow the packet header.
func (pp *PacketParser) parseLiteral(packet *Packet, bits []int8) (*Packet, []int8, error) {
	var literalBits []int8

	for len(bits) > 0 {
		split := bits[0:5]
		literalBits = append(literalBits, split[1:]...)

		bits = bits[5:]

		// break on last literal bit field
		if split[0] == 0 {
			break
		}
	}

	packet.SetValue(bitutils.Btoi(literalBits))

	return packet, bits, nil
}

func (pp *PacketParser) parseOperator(packet *Packet, bits []int8) (*Packet, []int8, error) {
	if packet.LengthTypeID() == TotalBitLength {
		totalLength := bitutils.Btoi(bits[:15])
		bits = bits[15:]

		if totalLength > len(bits) {
			totalLength = len(bits)
		}

		measuredBits := bits[:totalLength]

		for len(measuredBits) > 0 && bitutils.Btoi(measuredBits) != 0 {
			var err error
			_, measuredBits, err = pp.parsePacket(packet.AddChild(), measuredBits)
			if err != nil {
				return nil, nil, err
			}
		}

		packet.Evaluate()

		bits = bits[totalLength:]

		return packet, bits, nil
	}

	if packet.LengthTypeID() == NumSubPackets {
		n := bitutils.Btoi(bits[:11])
		bits = bits[11:]

		for i := 0; i < n; i++ {
			var err error
			_, bits, err = pp.parsePacket(packet.AddChild(), bits)
			if err != nil {
				return nil, nil, err
			}
		}

		packet.Evaluate()

		return packet, bits, nil
	}

	return nil, nil, errors.New("failed to parse operator")
}

func (pp *PacketParser) parsePacket(packet *Packet, bits []int8) (*Packet, []int8, error) {
	if bitutils.Btoi(bits) == 0 {
		return nil, nil, errors.New("bits zero valued")
	}

	packet, err := pp.parseHeader(packet, bits[:7])
	if err != nil {
		return nil, nil, err
	}

	switch packet.TypeID() {
	case Literal:
		_, bits, err = pp.parseLiteral(packet, bits[6:])
		if err != nil {
			return nil, nil, err
		}
	default:
		return pp.parseOperator(packet, bits[7:])
	}

	return packet, bits, nil
}

func (pp *PacketParser) Parse() (*Packet, error) {
	var bits []int8
	var err error

	rawArr := strings.Split(pp.raw, "")

	for _, char := range rawArr {
		bin, err := bitutils.Htob(char)
		if err != nil {
			return nil, err
		}

		bits = append(bits, bin...)
	}

	_, _, err = pp.parsePacket(pp.root, bits)
	if err != nil {
		return nil, err
	}

	pp.root.Evaluate()

	return pp.root, nil
}
