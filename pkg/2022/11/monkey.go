package day11

import (
	"math/big"
)

type MonkeyTest struct {
	DivisibleBy    *big.Int
	IfTrueThrowTo  int
	IfFalseThrowTo int
}

type MonkeyOperation struct {
	Op       string
	Constant *big.Int
}

type Monkey struct {
	Id              int
	StartingItems   []*big.Int
	Operation       MonkeyOperation
	Test            MonkeyTest
	InspectionCount int
}

func (m *Monkey) Inspect() *big.Int {
	item := m.StartingItems[0]
	m.StartingItems = m.StartingItems[1:]
	m.InspectionCount += 1

	return item
}

type OperationConstant func(*big.Int, *Monkey) *big.Int

func (m *Monkey) Operate(item, worryLevel *big.Int, fn OperationConstant) *big.Int {
	constant := fn(item, m)

	switch m.Operation.Op {
	case "+":
		item.Add(item, constant)
	case "-":
		item.Sub(item, constant)
	case "*":
		if item.Cmp(constant) == 0 {
			item.Exp(item, big.NewInt(2), nil)
		} else {
			item.Mul(item, constant)
		}
	case "/":
		item.Div(item, constant)
	}

	if worryLevel.Cmp(big.NewInt(1)) == 0 {
		return item
	}

	return item.Div(item, worryLevel)
}

func (m *Monkey) RunTest(val *big.Int) (monkeyId int) {
	z := new(big.Int)
	z.Mod(val, m.Test.DivisibleBy)
	if z.Cmp(big.NewInt(0)) == 0 {
		return m.Test.IfTrueThrowTo
	}

	return m.Test.IfFalseThrowTo
}
