package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	util "github.com/ralucas/advent-of-code/internal"
)

var inputFile = flag.String("input", "assets/8/input.txt", "Input file")

type Instruction struct {
	op   string
	sign string
	val  int
}

func prepareData(filepath string) []Instruction {
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
func runInstructions(instructions []Instruction) (int, int) {
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
func fixInstructions(instructions []Instruction) int {
	ilen := len(instructions)

	for i := 0; i < ilen; i++ {
		op := instructions[i].op
		if op == "nop" || op == "jmp" {
			cpInst := make([]Instruction, ilen)
			copy(cpInst, instructions)
			if op == "nop" {
				cpInst[i].op = "jmp"
			} else {
				cpInst[i].op = "nop"
			}

			acc, exitcode := runInstructions(cpInst)
			if exitcode == 0 {
				return acc
			}
		}

	}

	return -1
}

func main() {
	fmt.Print("Day 8\n===========\n")
	flag.Parse()
	data := prepareData(*inputFile)

	lastAcc, exitcode := runInstructions(data)
	if exitcode != -1 {
		log.Fatalf("Error, all instructions ran")
	}
	fmt.Println("A -- Last Accumulator:", lastAcc)

	finalAcc := fixInstructions(data)
	fmt.Println("B -- Final Accumulator:", finalAcc)
}
