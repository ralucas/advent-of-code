package day10

type InstructionType string

const (
	Noop InstructionType = "noop"
	Addx                 = "addx"
)

type Instruction struct {
	kind  InstructionType
	value int
}