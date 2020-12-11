package noop

type Day struct{}

func (d *Day) PrepareData(_ string) {
	return
}

func (d *Day) Part1() interface{} {
	return -1
}

func (d *Day) Part2() interface{} {
	return -1
}
