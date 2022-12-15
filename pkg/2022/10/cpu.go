package day10

type CPU struct {
	cycle          int
	registerValue  int
	signalStrength int
	store          Repository
}

type CPUState struct {
	cycle          int
	registerValue  int
	signalStrength int
}

type Repository interface {
	Record(CPUState) error
}

func NewCPU(repo Repository) *CPU {
	return &CPU{0, 1, 0, repo}
}

func (c *CPU) Eval(inst Instruction) error {
	switch inst.kind {
	case Noop:
		err := c.runCycle(0)
		if err != nil {
			return err
		}
	case Addx:
		err := c.runCycle(0)
		if err != nil {
			return err
		}

		err = c.runCycle(inst.value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CPU) runCycle(val int) error {
	c.cycle += 1
	c.signalStrength = c.cycle * c.registerValue
	c.registerValue += val

	err := c.store.Record(c.State())
	if err != nil {
		return err
	}

	return nil
}

func (c *CPU) State() CPUState {
	return CPUState{
		cycle:          c.cycle,
		registerValue:  c.registerValue,
		signalStrength: c.signalStrength,
	}
}
