package day11

import (
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	monkeys []*Monkey
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n\n")

	d.monkeys = make([]*Monkey, len(data))

	for _, m := range data {
		var err error

		monkey := &Monkey{}

		lines := strings.Split(m, "\n")

		for i, line := range lines {
			tline := strings.TrimSpace(line)
			sline := strings.Split(tline, ":")

			switch i {
			case 0:
				id, err := strconv.Atoi(strings.Split(sline[0], " ")[1])
				if err != nil {
					log.Fatal(err)
				}
				monkey.Id = id
			case 1:
				startingItems := make([]*big.Int, 0)

				items := strings.Split(sline[1], ",")
				for _, item := range items {
					iitem, err := strconv.Atoi(strings.TrimSpace(item))
					if err != nil {
						log.Fatal(err)
					}
					startingItems = append(startingItems, big.NewInt(int64(iitem)))
				}

				monkey.StartingItems = startingItems
			case 2:
				ss := strings.Split(sline[1], " = ")
				eq := strings.Split(ss[len(ss)-1], " ")
				op := eq[1]
				constant := eq[2]

				re := regexp.MustCompile(`^[0-9]+$`)

				iconstant := -1

				if re.MatchString(constant) {
					iconstant, err = strconv.Atoi(constant)
					if err != nil {
						log.Fatal(err)
					}
				}

				monkey.Operation = MonkeyOperation{Op: op, Constant: big.NewInt(int64(iconstant))}
			case 3:
				ss := strings.Split(sline[1], " ")
				divisibleBy, err := strconv.Atoi(ss[len(ss)-1])
				monkey.Test.DivisibleBy = big.NewInt(int64(divisibleBy))
				if err != nil {
					log.Fatal(err)
				}
			case 4:
				ss := strings.Split(sline[1], " ")
				monkey.Test.IfTrueThrowTo, err = strconv.Atoi(ss[len(ss)-1])
				if err != nil {
					log.Fatal(err)
				}
			case 5:
				ss := strings.Split(sline[1], " ")
				monkey.Test.IfFalseThrowTo, err = strconv.Atoi(ss[len(ss)-1])
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		d.monkeys[monkey.Id] = monkey
	}

	return
}

type PreOperateFunc func(item *big.Int, monkey *Monkey) *big.Int

func identity(item *big.Int, _ *Monkey) *big.Int {
	return item
}

// func popFn(item int, m *Monkey) int {
// 	if item < m.Test.DivisibleBy {
// 		return item
// 	}

// 	if m.Operation.Op != "*" {
// 		return item
// 	}

// 	rem := item % m.Test.DivisibleBy
// 	return m.Test.DivisibleBy + rem
// }

// func mod(item int, _ *Monkey) int {
// 	return item % 100000000
// }

// func onlyDbl(item int, m *Monkey) int {
// 	if m.Operation.Constant == -1 {
// 		rem := item % m.Test.DivisibleBy
// 		return m.Test.DivisibleBy + rem
// 	}

// 	return item
// }

func stdOperationConstant(item *big.Int, m *Monkey) *big.Int {
	if m.Operation.Constant.Cmp(big.NewInt(-1)) == 0 {
		return item
	}

	return m.Operation.Constant
}

// func remOpConst(item int, m *Monkey) int {
// 	rem := item % m.Test.DivisibleBy

// 	if m.Operation.Constant == -1 {
// 		return m.Test.DivisibleBy + rem
// 	}

// 	return rem
// }

// func dblOpConst(item int, m *Monkey) int {
// 	if m.Operation.Constant == -1 {
// 		return 1
// 	}

// 	return m.Operation.Constant
// }

func (d *Day) RunRound(monkeyId int, worryLevel *big.Int, _ PreOperateFunc, opConstFn OperationConstant) error {
	monkey := d.monkeys[monkeyId]

	mlen := len(monkey.StartingItems)

	var wg sync.WaitGroup

	mu := &sync.Mutex{}

	for i := 0; i < mlen; i++ {
		wg.Add(1)
		v := monkey.Inspect()
		go func(item *big.Int, m *Monkey, wlvl *big.Int, fn OperationConstant) {
			// item = fn(item, monkey)
			defer wg.Done()

			newVal := m.Operate(item, wlvl, fn)

			// if newVal < 0 {
			// 	return fmt.Errorf("overflow detected for id [%d] %d", monkeyId, newVal)
			// }

			throwToMonkeyId := m.RunTest(newVal)

			throwToMonkey := d.monkeys[throwToMonkeyId]

			mu.Lock()
			throwToMonkey.StartingItems = append(throwToMonkey.StartingItems, newVal)

			d.monkeys[throwToMonkeyId] = throwToMonkey
			mu.Unlock()
		}(v, monkey, worryLevel, opConstFn)
	}

	wg.Wait()
	d.monkeys[monkeyId] = monkey

	return nil
}

func (d *Day) Part1() interface{} {
	rounds := 20

	worryLevel := big.NewInt(int64(3))

	for round := 0; round < rounds; round++ {
		for id := range d.monkeys {
			d.RunRound(id, worryLevel, identity, stdOperationConstant)
		}
	}

	inspections := make([]int, 0)
	for _, monkey := range d.monkeys {
		inspections = append(inspections, monkey.InspectionCount)
	}

	sorted := sortutil.QSort(inspections)

	return sorted[len(sorted)-1] * sorted[len(sorted)-2]
}

func (d *Day) Part2() interface{} {
	rounds := 800

	worryLevel := big.NewInt(int64(1))

	for round := 0; round < rounds; round++ {
		for id := range d.monkeys {
			// d.printRoundState(round + 1)
			err := d.RunRound(id, worryLevel, identity, stdOperationConstant)
			if err != nil {
				log.Fatalf("round %d: %+v", round, err)
			}
		}
		fmt.Printf("completed round: %d\n", round)
		// d.printRoundState(round + 1)
	}

	inspections := make([]int, 0)
	for _, monkey := range d.monkeys {
		inspections = append(inspections, monkey.InspectionCount)
	}

	sorted := sortutil.QSort(inspections)

	return sorted[len(sorted)-1] * sorted[len(sorted)-2]
}

func (d *Day) printRoundState(round int) {
	fmt.Printf("-- Round %d --\n", round)
	for _, m := range d.monkeys {
		fmt.Printf("ID %d, inspectCt %d, items %d %+v\n", m.Id, m.InspectionCount, len(m.StartingItems), m.StartingItems)
	}
}
