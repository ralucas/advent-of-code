package day8

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/ralucas/advent-of-code/pkg/util"
)

var inputFile = flag.String("input", "assets/8/input.txt", "Input file")

type Instruction struct {
	op   string
	sign string
	val  int
}

func PrepareData(filepath string) []Instruction {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := util.ReadFileToArray(filepath, "\n")

	var instructions []Instruction

	for _, d := range data {
		spl := strings.Split(d, " ")
		val, err := strconv.Atoi(spl[1][1:])
		if err != nil {
			log.Fatalf("Error parsing instruction value %+v\n", err)
		}
		inst := Instruction{
			op:   spl[0],
			sign: string(spl[1][0]),
			val:  val,
		}
		instructions = append(instructions, inst)
	}

	return instructions
}

// runInstructions returns last good accumulator and exit code
// It runs all input instructions only breaking prior to
// replaying instructions
func RunInstructions(instructions []Instruction) (int, int) {
	acc := 0

	visited := make([]int, len(instructions))
	exitcode := 0

	i := 0
	ilen := len(instructions)

	for i < ilen {
		if visited[i] == 1 {
			return acc, -1
		}

		visited[i] = 1

		instruction := instructions[i]

		switch instruction.op {
		case "acc":
			if instruction.sign == "+" {
				acc += instruction.val
			} else {
				acc -= instruction.val
			}
			i++
		case "jmp":
			if instruction.sign == "+" {
				i += instruction.val
			} else {
				i -= instruction.val
			}
		case "nop":
			i++
		}
	}

	return acc, exitcode
}

// fixInstructions
// runtime O(n^2)
func FixInstructions(instructions []Instruction) int {
	ilen := len(instructions)

	ch := make(chan int, ilen)

	for i := 0; i < ilen; i++ {
		operation := instructions[i].op

		go func(idx int, op string, insts []Instruction) {
			if op == "nop" || op == "jmp" {
				// do copy here inside if stmt
				cp := make([]Instruction, len(insts))
				copy(cp, insts)

				if op == "nop" {
					cp[idx].op = "jmp"
				} else {
					cp[idx].op = "nop"
				}

				acc, exitcode := RunInstructions(cp)
				if exitcode == 0 {
					ch <- acc
				} else {
					ch <- -1
				}
			} else {
				ch <- -1
			}
		}(i, operation, instructions)
	}

	for i := 0; i < ilen; i++ {
		acc := <-ch
		if acc != -1 {
			return acc
		}
	}

	defer close(ch)
	return -1
}
