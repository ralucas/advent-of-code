package day4

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Passport struct {
	byr    int
	iyr    int
	eyr    int
	hgt    int
	huom   string
	og_hgt string
	hcl    string
	ecl    string
	pid    string
	cid    string
}

type Day struct {
	data []Passport
}

func (d *Day) PrepareData(filepath string) {
	data := fileutil.ReadFileToArray(filepath, "\n\n")

	var preparedData []Passport

	for _, line := range data {
		l1 := strings.Split(line, "\n")
		var l2 []string
		for _, l := range l1 {
			l2 = append(l2, strings.Split(l, " ")...)
		}
		var l3 []string
		for _, l := range l2 {
			l3 = append(l3, strings.Split(l, ":")...)
		}
		pass := Passport{}
		re := regexp.MustCompile(`[0-9]+`)
		var err error
		for i, item := range l3 {
			switch item {
			case "byr":
				pass.byr, err = strconv.Atoi(strings.TrimSpace(l3[i+1]))
				if err != nil {
					log.Fatalf("%v", err)
				}
			case "iyr":
				pass.iyr, err = strconv.Atoi(strings.TrimSpace(l3[i+1]))
				if err != nil {
					log.Fatalf("%v", err)
				}
			case "eyr":
				pass.eyr, err = strconv.Atoi(strings.TrimSpace(l3[i+1]))
				if err != nil {
					log.Fatalf("%v", err)
				}
			case "hgt":
				pass.og_hgt = l3[i+1]
				hgt := re.FindStringIndex(l3[i+1])
				pass.hgt, err = strconv.Atoi(l3[i+1][hgt[0]:hgt[1]])
				if err != nil {
					log.Fatalf("%v", err)
				}
				pass.huom = l3[i+1][hgt[1]:]
			case "hcl":
				pass.hcl = strings.TrimSpace(l3[i+1])
			case "ecl":
				pass.ecl = strings.TrimSpace(l3[i+1])
			case "pid":
				pass.pid = strings.TrimSpace(l3[i+1])
			case "cid":
				pass.cid = strings.TrimSpace(l3[i+1])
			}
		}
		preparedData = append(preparedData, pass)
	}

	d.data = preparedData

	return
}

func (d *Day) Part1() interface{} {
	validPassportCount := CountValidPassports(d.data)

	return validPassportCount
}

func (d *Day) Part2() interface{} {
	validPassportStrictCount := CountValidPassportsStrict(d.data)

	return validPassportStrictCount
}

func isValid(p Passport) bool {
	return p.byr != 0 &&
		p.iyr != 0 &&
		p.eyr != 0 &&
		p.hgt != 0 &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

// isValidStrict validates a passport on the following rules:
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
//
//	If cm, the number must be at least 150 and at most 193.
//	If in, the number must be at least 59 and at most 76.
//
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func isValidStrict(p Passport) bool {
	validByr := p.byr >= 1920 && p.byr <= 2002
	validIyr := p.iyr >= 2010 && p.iyr <= 2020
	validEyr := p.eyr >= 2020 && p.eyr <= 2030
	validHgt := isValidHgt(p.hgt, p.huom)
	validHcl := isValidHcl(p.hcl)
	validEcl := isValidEcl(p.ecl)
	validPid := isValidPid(p.pid)

	result := validByr && validIyr && validEyr && validHgt &&
		validHcl && validEcl && validPid

	return result
}

// isValidHgt evaluates on
// hgt (Height) - a number followed by either cm or in:
//
//	If cm, the number must be at least 150 and at most 193.
//	If in, the number must be at least 59 and at most 76.
func isValidHgt(hgt int, uom string) bool {
	validHgt := false
	if uom == "cm" {
		validHgt = hgt >= 150 && hgt <= 193
	}
	if uom == "in" {
		validHgt = hgt >= 59 && hgt <= 76
	}

	return validHgt
}

// isValidHcl evaluates on:
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func isValidHcl(hcl string) bool {
	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	return len(hcl) == 7 && re.Match([]byte(hcl))
}

// isValidEcl
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func isValidEcl(ecl string) bool {
	possibilities := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, possible := range possibilities {
		if ecl == possible {
			return true
		}
	}
	return false
}

// isValidPid
// pid (Passport ID) - a nine-digit number, including leading zeroes.
func isValidPid(pid string) bool {
	re := regexp.MustCompile(`[0-9]{9}`)
	return len(pid) == 9 && re.Match([]byte(pid))
}

func CountValidPassports(vp []Passport) int {
	count := 0
	for _, p := range vp {
		if isValid(p) {
			count += 1
		}
	}
	return count
}

func CountValidPassportsStrict(vp []Passport) int {
	count := 0
	for _, p := range vp {
		if isValid(p) && isValidStrict(p) {
			count += 1
		}
	}
	return count
}
