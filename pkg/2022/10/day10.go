package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	// bitutil "github.com/ralucas/advent-of-code/pkg/util/bit"
	// mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
	// sortutil "github.com/ralucas/advent-of-code/pkg/util/sort"
)

type Day struct {
	instructions []Instruction
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	d.instructions = make([]Instruction, len(data))
	for i, s := range data {
		ss := strings.Split(s, " ")
		instKind := InstructionType(ss[0])

		switch instKind {
		case Noop:
			d.instructions[i] = Instruction{Noop, 0}
		case Addx:
			val, err := strconv.Atoi(ss[1])
			if err != nil {
				log.Fatal(err)
			}
			d.instructions[i] = Instruction{Addx, val}
		}
	}

	return
}

func (d *Day) Part1() interface{} {
	recorder := NewRecorder()
	cpu := NewCPU(recorder)

	for _, instruction := range d.instructions {
		cpu.Eval(instruction)
	}

	total := 0
	for i := 20; i <= 220; i += 40 {
		state, err := recorder.RetrieveByCycle(i)
		if err != nil {
			log.Fatal(err)
		}
		total += state.signalStrength
	}

	return total
}

func (d *Day) Part2() interface{} {
	// run instructions
	recorder := NewRecorder()
	cpu := NewCPU(recorder)

	for _, instruction := range d.instructions {
		cpu.Eval(instruction)
	}

	// create CRT
	crt := make([][]byte, 0)
	size := recorder.Size()
	for i := 1; i < size; i += 40 {
		b := make([]byte, 0)
		for j := 0; j < 40; j++ {
			b = append(b, []byte(".")...)
		}
		crt = append(crt, b)
	}

	// draw pixels
	pixelMin, pixelMax := 0, 2

	lit := []byte("#")

	for i, row := range crt {
		start := i*40 + 1

		for j := 0; j < 40; j++ {
			cycle := start + j
			cpuState, err := recorder.RetrieveByCycle(cycle)
			if err != nil {
				log.Fatal(err)
			}
			register := cpuState.registerValue

			if j >= pixelMin && j <= pixelMax {
				row[j] = lit[0]
			}

			pixelMin, pixelMax = register-1, register+1

		}
	}

	sCrt := make([]string, len(crt))

	for i := range crt {
		sCrt[i] = string(crt[i])
		fmt.Println(sCrt[i])
	}

	return sCrt
}
