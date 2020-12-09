package day2

import (
	"strings"

	"github.com/ralucas/advent-of-code/pkg/util"
)

type Password struct {
	min      int
	max      int
	letter   string
	password string
}

func PrepareData(filepath string) []Password {
	data := util.ReadFile(filepath)
	dataArr := strings.Split(data, "\n")
	output := make([]Password, 0)
	for _, line := range dataArr {
		if line != "" {
			splitLine := strings.Split(line, " ")
			minmax := util.MapToInt(strings.Split(splitLine[0], "-"))
			letter := string(splitLine[1][0])
			password := splitLine[len(splitLine)-1]
			newPw := Password{minmax[0], minmax[1], letter, password}
			output = append(output, newPw)
		}
	}
	return output
}

func IsValid(p Password) bool {
	matchCount := 0
	for i := 0; i < len(p.password); i++ {
		if string(p.password[i]) == p.letter {
			matchCount += 1
		}
	}

	return matchCount >= p.min && matchCount <= p.max
}

func IsValidByPosition(p Password) bool {
	pos1 := p.min - 1
	pos2 := p.max - 1

	valid1 := string(p.password[pos1]) == p.letter && string(p.password[pos2]) != p.letter
	valid2 := string(p.password[pos1]) != p.letter && string(p.password[pos2]) == p.letter

	return valid1 || valid2
}
