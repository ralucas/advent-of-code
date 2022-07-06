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
	packets      []*Packet
	lengthTypeID LengthTypeID
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

	if packet.TypeID() == Operator {
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
func (pp *PacketParser) parseLiteral(packet *Packet, bits []int8) error {
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

	packet.literal = bitutils.Btoi(literalBits)

	if len(bits) == 0 || bitutils.Btoi(bits) == 0 {
		return nil
	}

	return pp.parsePacket(bits)
}

func (pp *PacketParser) parseOperator(packet *Packet, bits []int8) error {
	if packet.LengthTypeID() == TotalBitLength {
		totalLength := bitutils.Btoi(bits[:15])
		bits = bits[15:]

		if totalLength > len(bits) {
			totalLength = len(bits)
		}

		return pp.parsePacket(bits[:totalLength])
	}

	if packet.LengthTypeID() == NumSubPackets {
		n := bitutils.Btoi(bits[:11])
		packet.numChildren = n	
		bits = bits[11:]
		// bits = trimLeadingZeros(bits)
		bitslen := len(bits) / n

		for i := 0; i < n; i++ {
			start, end := i*bitslen, (i+1)*bitslen

			err := pp.parsePacket(bits[start:end])
			if err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New("failed to parse operator")
}

func (pp *PacketParser) parsePacket(bits []int8) error {
	// create new packet on each call
	packet := pp.newPacket()

	packet, err := pp.parseHeader(packet, bits[:7])
	if err != nil {
		return err
	}

	if packet.TypeID() == Literal {
		return pp.parseLiteral(packet, bits[6:])
	}

	if packet.TypeID() == Operator {
		return pp.parseOperator(packet, bits[7:])
	}

	return errors.New("failed to parse packet")
}

func (pp *PacketParser) newPacket() *Packet {
	packet := new(Packet)
	pp.packets = append(pp.packets, packet)

	return packet
}

func (pp *PacketParser) Parse() ([]*Packet, error) {
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

	// bits = trimTrailingZeros(bits)

	err = pp.parsePacket(bits)
	if err != nil {
		return nil, err
	}

	return pp.packets, nil
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
