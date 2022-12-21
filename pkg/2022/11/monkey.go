package day11

type MonkeyTest struct {
	DivisibleBy    int
	IfTrueThrowTo  int
	IfFalseThrowTo int
}

type MonkeyOperation struct {
	Op       string
	Constant int
}

type Monkey struct {
	Id              int
	StartingItems   []int
	Operation       MonkeyOperation
	Test            MonkeyTest
	InspectionCount int
}

func (m *Monkey) Inspect() int {
	item := m.StartingItems[0]
	m.StartingItems = m.StartingItems[1:]
	m.InspectionCount += 1

	return item
}

type OperationConstant func(int, *Monkey) int

func (m *Monkey) Operate(item, worryLevel int, fn OperationConstant) int {
	ans := 0

	constant := fn(item, m)

	switch m.Operation.Op {
	case "+":
		ans = item + constant
	case "-":
		ans = item - constant
	case "*":
		ans = item * constant
	case "/":
		ans = item / constant
	}

	if worryLevel == 1 {
		return ans
	}

	return ans / worryLevel
}

func (m *Monkey) RunTest(val int) (monkeyId int) {
	if val%m.Test.DivisibleBy == 0 {
		return m.Test.IfTrueThrowTo
	}

	return m.Test.IfFalseThrowTo
}
