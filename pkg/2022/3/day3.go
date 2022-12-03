package day3

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	data []string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	d.data = data

	return
}

func (d *Day) Part1() interface{} {
	total := 0

	for _, chars := range d.data {
		l, r := 0, len(chars)-1

		ml := make(map[byte]bool)
		mr := make(map[byte]bool)

		for l <= r {
			cl := chars[l]
			cr := chars[r]

			// handle edge case where they equal
			// another at the same time
			if cl == cr {
				total += priority(cr)
				break
			}

			if ok, _ := ml[cr]; ok {
				total += priority(cr)
				break
			}
			if ok, _ := mr[cl]; ok {
				total += priority(cl)
				break
			}
			mr[cr] = true
			ml[cl] = true
			l++
			r--
		}
	}

	return total
}

func priority(char byte) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	}

	return int(char-'A') + 27
}

func (d *Day) Part2() interface{} {
	total := 0

	for i := 0; i <= len(d.data)-3; i += 3 {
		group := d.data[i : i+3]
		m := make(map[byte]int)

		for idx, chars := range group {
			for j := range chars {
				char := chars[j]
				v, ok := m[char]
				if (ok && v == idx) || (idx == 0) {
					m[char] = idx + 1
				}
				if m[char] == 3 {
					total += priority(char)
					break
				}
			}
		}
	}
	return total
}
