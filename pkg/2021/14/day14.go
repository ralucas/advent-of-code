package day14

import (
	"log"
	"math"
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
)

type Day struct {
	polymer      string
	rules        map[string]string
	letterCounts map[string]int
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutils.ReadFile(filepath)

	spl := strings.Split(data, "\n\n")

	d.polymer = spl[0]

	d.letterCounts = make(map[string]int)

	for i := range d.polymer {
		d.letterCounts[string(d.polymer[i])] += 1
	}

	ruleLines := strings.Split(spl[1], "\n")

	d.rules = make(map[string]string)

	for _, line := range ruleLines {
		ruleSpl := strings.Split(line, " -> ")
		d.rules[ruleSpl[0]] = ruleSpl[1]
	}

	return
}

func (d *Day) insert(s string) string {
	var sb strings.Builder

	var sbTotal strings.Builder

	for i := 1; i < len(s); i++ {
		pair := s[i-1 : i+1]
		if v, ok := d.rules[pair]; ok {
			sb.Reset()
			sb.WriteByte(pair[0])
			sb.WriteString(v)
		}
		sbTotal.WriteString(sb.String())
	}

	sbTotal.WriteByte(s[len(s)-1])
	return sbTotal.String()
}

func (d *Day) Part1() interface{} {
	s := d.polymer
	for i := 0; i < 10; i++ {
		s = d.insert(s)
	}

	elementCounts := make(map[rune]int)
	for _, r := range s {
		elementCounts[r] += 1
	}

	maxE, minE := 0, math.MaxInt64
	for _, v := range elementCounts {
		if v > maxE {
			maxE = v
		}

		if v < minE {
			minE = v
		}
	}

	return maxE - minE
}

func (d *Day) pairs(start map[string]int) map[string]int {
	m := make(map[string]int)

	var sb strings.Builder

	for k, v := range start {
		nk := d.rules[k]
		// handle 1st
		sb.WriteByte(k[0])
		sb.WriteString(nk)
		m[sb.String()] += v
		sb.Reset()

		// handle 2nd
		sb.WriteString(nk)
		sb.WriteByte(k[1])
		m[sb.String()] += v
		sb.Reset()
	}

	return m
}

func (d *Day) sToMap(s string) map[string]int {
	m := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		m[s[i:i+2]] += 1
	}

	return m
}

func (d *Day) Part2() interface{} {
	m := d.sToMap(d.polymer)

	for i := 0; i < 40; i++ {
		m = d.pairs(m)
	}

	lc := make(map[string]int)

	for k, v := range m {
		lc[string(k[0])] += v
		lc[string(k[1])] += v
	}

	maxE, minE := 0, math.MaxInt

	for _, v := range lc {
		if v > maxE {
			maxE = v
		}

		if v < minE {
			minE = v
		}
	}

	return int(math.Round((float64(maxE) / 2.0)) - math.Round((float64(minE) / 2.0)))
}
