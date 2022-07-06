package day2

import (
	"strings"

	fileutils "github.com/ralucas/advent-of-code/pkg/utils/file"
	arrayutils "github.com/ralucas/advent-of-code/pkg/utils/array"
)

type Password struct {
	Min    int
	Max    int
	Letter string
	Pass   string
}

type Day struct {
	data []Password
}

func (d *Day) PrepareData(filepath string) {
	data := fileutils.ReadFile(filepath)
	dataArr := strings.Split(data, "\n")
	output := make([]Password, 0)
	for _, line := range dataArr {
		if line != "" {
			splitLine := strings.Split(line, " ")
			minmax := arrayutils.MapToInt(strings.Split(splitLine[0], "-"))
			letter := string(splitLine[1][0])
			password := splitLine[len(splitLine)-1]
			newPw := Password{minmax[0], minmax[1], letter, password}
			output = append(output, newPw)
		}
	}

	d.data = output

	return
}

func (d *Day) Part1() interface{} {
	validCount := 0
	for _, pw := range d.data {
		if IsValid(pw) {
			validCount += 1
		}
	}

	return validCount
}

func (d *Day) Part2() interface{} {
	validPosCount := 0
	for _, pw := range d.data {
		if IsValidByPosition(pw) {
			validPosCount += 1
		}
	}

	return validPosCount
}

func IsValid(p Password) bool {
	matchCount := 0
	for i := 0; i < len(p.Pass); i++ {
		if string(p.Pass[i]) == p.Letter {
			matchCount += 1
		}
	}

	return matchCount >= p.Min && matchCount <= p.Max
}

func IsValidByPosition(p Password) bool {
	pos1 := p.Min - 1
	pos2 := p.Max - 1

	valid1 := string(p.Pass[pos1]) == p.Letter && string(p.Pass[pos2]) != p.Letter
	valid2 := string(p.Pass[pos1]) != p.Letter && string(p.Pass[pos2]) == p.Letter

	return valid1 || valid2
}
