package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareData(t *testing.T) {
	d := Day{}
	d.PrepareData("../../../assets/2020/4/input.txt")

	t.Run("is correct type", func(t *testing.T) {
		assert.IsType(t, d.data[0], Passport{})
	})

	t.Run("has items", func(t *testing.T) {
		assert.Equal(t, 280, len(d.data))
	})
}

func TestIsValidPassport(t *testing.T) {
	t.Run("is not a valid passport", func(t *testing.T) {
		testPassports := []Passport{
			{pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13"},
		}
		for _, testPass := range testPassports {
			assert.False(t, isValid(testPass))
		}
	})

	t.Run("is a valid passport", func(t *testing.T) {
		testPassports := []Passport{
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
		}
		for _, testPass := range testPassports {
			assert.True(t, isValid(testPass))
		}
	})
}

func TestIsValidHgt(t *testing.T) {
	type test struct {
		hgt int
		uom string
	}
	t.Run("invalid hgt", func(t *testing.T) {
		testHgts := []test{
			{hgt: 181},
			{uom: "cm"},
			{hgt: 194, uom: "cm"},
			{hgt: 149, uom: "cm"},
			{hgt: 49, uom: "in"},
			{hgt: 80, uom: "in"},
		}
		for _, testHgt := range testHgts {
			assert.False(t, isValidHgt(testHgt.hgt, testHgt.uom))
		}
	})

	t.Run("valid hgt", func(t *testing.T) {
		testHgts := []test{
			{hgt: 193, uom: "cm"},
			{hgt: 150, uom: "cm"},
			{hgt: 175, uom: "cm"},
			{hgt: 59, uom: "in"},
			{hgt: 76, uom: "in"},
			{hgt: 65, uom: "in"},
		}
		for _, testHgt := range testHgts {
			assert.True(t, isValidHgt(testHgt.hgt, testHgt.uom))
		}
	})
}

func TestIsValidEcl(t *testing.T) {
	t.Run("invalid ecl", func(t *testing.T) {
		tests := []string{"", "has", "ame"}
		for _, test := range tests {
			assert.False(t, isValidEcl(test))
		}
	})

	t.Run("valid ecl", func(t *testing.T) {
		tests := []string{
			"hzl",
			"amb",
			"blu",
			"brn",
			"gry",
			"grn",
			"oth",
		}
		for _, test := range tests {
			assert.True(t, isValidEcl(test))
		}
	})
}

func TestIsValidHcl(t *testing.T) {
	t.Run("invalid hcl", func(t *testing.T) {
		tests := []string{
			"341e138",
			"#341e1",
			"#341g13",
			"",
			"#341",
		}
		for _, test := range tests {
			assert.False(t, isValidHcl(test))
		}
	})

	t.Run("valid hcl", func(t *testing.T) {
		tests := []string{
			"#341e13",
			"#ffffff",
			"#111111",
			"#a41e13",
		}
		for _, test := range tests {
			assert.True(t, isValidHcl(test))
		}
	})
}

func TestIsValidPid(t *testing.T) {
	t.Run("invalid pid", func(t *testing.T) {
		tests := []string{
			"23705680",
			"",
			"has",
			"0",
			"00000000",
			"12345678a",
			"6000619833",
		}
		for _, test := range tests {
			assert.False(t, isValidPid(test))
		}
	})

	t.Run("valid pid", func(t *testing.T) {
		tests := []string{
			"623705680",
			"623705681",
			"000005680",
		}
		for _, test := range tests {
			assert.True(t, isValidPid(test))
		}
	})
}

func TestIsValidPassportStrict(t *testing.T) {
	t.Run("is not a valid passport", func(t *testing.T) {
		testPassports := []Passport{
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1919, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2009, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 2003, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2021, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2019, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2031, cid: "test"},
			{iyr: 2010, pid: "23705680", hgt: 193, huom: "cm", byr: 1980, hcl: "341e138", eyr: 2028},
			{iyr: 2010, ecl: "has", hgt: 193, huom: "cm", byr: 1980, hcl: "#341e1", eyr: 2028},
			{iyr: 2010, pid: "0", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, hcl: "#341g13", eyr: 2028},
			{iyr: 2010, pid: "00000000", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, eyr: 2028},
			{iyr: 2010, pid: "12345678a", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, hcl: "#341", eyr: 2028},
			{iyr: 2010, pid: "623705680", hgt: 193, huom: "cm", byr: 1980, hcl: "341e138", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "has", hgt: 193, huom: "cm", byr: 1980, hcl: "#341e1", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, hcl: "#341g13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, hcl: "#341", eyr: 2028},
			{iyr: 2010, pid: "623705680", hgt: 193, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "has", hgt: 193, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "ame", hgt: 193, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 194, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 149, huom: "cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 49, huom: "in", byr: 1980, hcl: "#341e13", eyr: 2028},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 80, huom: "in", byr: 1980, hcl: "#341e13", eyr: 2028},
			{eyr: 1972, cid: "100", hcl: "#18171d", ecl: "amb", hgt: 170, pid: "186cm", iyr: 2018, byr: 1926},
			{iyr: 2019, hcl: "#602927", eyr: 1967, hgt: 170, huom: "cm", ecl: "grn", pid: "012533040", byr: 1946},
			{hcl: "dab227", iyr: 2012, ecl: "brn", hgt: 182, huom: "cm", pid: "021572410", eyr: 2020, byr: 1992, cid: "277"},
			{hgt: 59, huom: "cm", ecl: "zzz", eyr: 2038, hcl: "74454a", iyr: 2023, pid: "3556412378", byr: 2007},
		}
		for _, testPass := range testPassports {
			assert.False(t, isValidStrict(testPass))
		}
	})

	t.Run("is a valid passport", func(t *testing.T) {
		testPassports := []Passport{
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", og_hgt: "181cm", byr: 1980, hcl: "#341e13", eyr: 2028, cid: "test"},
			{iyr: 2010, pid: "623705680", ecl: "hzl", hgt: 181, huom: "cm", og_hgt: "181cm", byr: 1980, hcl: "#341e13", eyr: 2028},
			{pid: "087499704", hgt: 74, huom: "in", og_hgt: "74in", ecl: "grn", iyr: 2012, eyr: 2030, byr: 1980, hcl: "#623a2f"},
			{eyr: 2029, ecl: "blu", cid: "129", byr: 1989, iyr: 2014, pid: "896056539", hcl: "#a97842", hgt: 165, huom: "cm", og_hgt: "165cm"},
			{hcl: "#888785", hgt: 164, huom: "cm", og_hgt: "164cm", byr: 2001, iyr: 2015, cid: "88", pid: "545766238", ecl: "hzl", eyr: 2022},
			{iyr: 2010, hgt: 158, huom: "cm", og_hgt: "158cm", hcl: "#b6652a", ecl: "blu", byr: 1944, eyr: 2021, pid: "093154719"},
		}
		for _, testPass := range testPassports {
			assert.True(t, isValidStrict(testPass), fmt.Sprintf("Failing Passport: %+v\n", testPass))
		}
	})
}

func TestIntegrationTotalValidity(t *testing.T) {
	t.Run("is a not valid passport from file", func(t *testing.T) {
		d := Day{}
		d.PrepareData("../../../test/testdata/2020/4/invalid.txt")
		for _, p := range d.data {
			valid := isValid(p) && isValidStrict(p)
			assert.False(t, valid)
		}
	})

	t.Run("is a valid passport from file", func(t *testing.T) {
		d := Day{}
		d.PrepareData("../../../test/testdata/2020/4/valid.txt")
		for _, p := range d.data {
			valid := isValid(p) && isValidStrict(p)
			assert.True(t, valid)
		}
	})
}

func TestCountValidPassports(t *testing.T) {
	d := Day{}
	d.PrepareData("../../../test/testdata/2020/4/test_input_a.txt")
	assert.Equal(t, 2, CountValidPassports(d.data))
}

func TestCountValidPassportsStrict(t *testing.T) {
	d := Day{}
	d.PrepareData("../../../test/testdata/2020/4/test_input_b.txt")
	assert.Equal(t, 4, CountValidPassportsStrict(d.data))
}
