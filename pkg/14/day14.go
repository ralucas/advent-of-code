package day14

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	data []ProgramSet
}

type ProgramSet struct {
	mask string
	mem  []Memory
}

type Memory struct {
	loc    int
	val    int
	valbin []int8
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	var p ProgramSet
	re := regexp.MustCompile(`\d+`)
	for i, line := range data {
		l := strings.Split(line, " = ")
		if l[0] == "mask" {
			if i > 0 {
				d.data = append(d.data, p)
			}
			p = ProgramSet{mask: l[1]}
		} else {
			loc, err := strconv.Atoi(string(re.Find([]byte(l[0]))))
			if err != nil {
				log.Fatalf("error strconv %v", err)
			}
			val, err := strconv.Atoi(l[1])
			if val > (1<<36)-1 {
				log.Fatalf("TOO BIG: %d", val)
			}
			if err != nil {
				log.Fatalf("error strconv %v", err)
			}
			mem := NewMemory(loc, val)
			p.mem = append(p.mem, mem)
		}
	}
	d.data = append(d.data, p)

	return
}

func (d *Day) Part1() interface{} {
	mmap := make(map[int]int)

	for _, set := range d.data {
		mask := set.mask
		for _, m := range set.mem {
			loc := m.loc
			am := ApplyMask(m.valbin, mask)
			mmap[loc] = btoi(am)
		}
	}

	output := int64(0)

	for _, v := range mmap {
		output += int64(v)
	}

	return output
}

func (d *Day) Part2() interface{} {
	return -1
}

func NewMemory(loc int, val int) Memory {
	return Memory{
		loc:    loc,
		val:    val,
		valbin: itob(val),
	}
}

func itob(val int) []int8 {
	bin := make([]int8, 36)
	pow := 1
	i := 0
	for i < 36 {
		if val>>i == 1 {
			bin[35-i] = 1
			val -= pow

			if val == 0 {
				break
			}
			// reset
			i = 0
			pow = 1
			continue
		}

		pow *= 2
		i++
	}

	return bin
}

func btoi(b []int8) int {
	r := len(b) - 1

	output := 0
	rs := 1

	for r >= 0 {
		if b[r] == 1 {
			output += rs
		}
		r--
		rs *= 2
	}

	return output
}

func ApplyMask(bin []int8, mask string) []int8 {
	for i, r := range mask {
		switch r {
		case rune('1'):
			if bin[i] == 0 {
				bin[i] = 1
			}
		case rune('0'):
			if bin[i] == 1 {
				bin[i] = 0
			}
		}
	}

	return bin
}
