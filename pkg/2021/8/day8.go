package day8

import (
	"log"
	"strconv"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	// TODO: Change this
	signalPatterns [][]string
	outputValues   [][]string
}

var numsToLetters = []string{
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	for _, line := range data {
		spl := strings.Split(line, "|")
		sp, ov := spl[0], spl[1]
		d.signalPatterns = append(d.signalPatterns, arrayutil.Map(
			strings.Split(sp, " "), func(s string, i int) string {
				return strings.TrimSpace(s)
			}),
		)
		d.outputValues = append(d.outputValues, arrayutil.Map(
			strings.Split(ov, " "), func(s string, i int) string {
				return strings.TrimSpace(s)
			}),
		)
	}

	return
}

func (d *Day) Part1() interface{} {
	count := 0
	for _, ovs := range d.outputValues {
		f := arrayutil.Filter(ovs, func(s string) bool {
			return len(s) == 2 ||
				len(s) == 3 ||
				len(s) == 4 ||
				len(s) == 7
		})

		count += len(f)
	}

	return count
}

func (d *Day) Part2() interface{} {
	// lens := arrayutil.MapToInt2(numsToLetters, func(s string, i int) int {
	// 	return len(s)
	// })
	var total int

	for i, sps := range d.signalPatterns {
		nums := make([]string, 10)
		n235 := make([]string, 0)
		n069 := make([]string, 0)

		for _, sp := range sps {
			switch len(sp) {
			case 7:
				nums[8] = sp
			case 2:
				nums[1] = sp
			case 3:
				nums[7] = sp
			case 4:
				nums[4] = sp
			case 5:
				n235 = append(n235, sp)
			case 6:
				n069 = append(n069, sp)
			}
		}

		// 3: has both letters from 1
		// 5: subtract letters of 1 from 5, 2 and 4;
		//   5 has all the letters that 4 now has
		// 2: is what's left
		for _, n := range n235 {
			if isThree(nums[1], n) {
				nums[3] = n
				continue
			}

			if isFive(nums[1], nums[4], n) {
				nums[5] = n
				continue
			}

			nums[2] = n
		}

		for _, n := range n069 {
			if isSix(nums[1], n) {
				nums[6] = n
				continue
			}

			if isNine(nums[4], n) {
				nums[9] = n
				continue
			}

			nums[0] = n
		}

		lineNum := ""
		for _, ov := range d.outputValues[i] {
			for x, n := range nums {
				if len(n) == len(ov) {
					if isSame(n, ov) {
						lineNum += strconv.Itoa(x)
					}
				}
			}
		}

		t, _ := strconv.Atoi(lineNum)
		total += t
	}

	return total
}

func isSame(a, b string) bool {
	v := make([]int, 7)

	for _, r := range a {
		idx := r - 97
		v[idx] += 1
	}

	for _, r := range b {
		idx := r - 97
		if v[idx] == 0 {
			return false
		}
	}

	return true
}

func isThree(one, s string) bool {
	count := 0

	for _, b := range []byte(s) {
		if b == one[0] || b == one[1] {
			count++
		}
	}

	return count == 2
}

func isFive(one, four, s string) bool {
	fourX := make([]rune, 0)
	for _, b := range four {
		if b != rune(one[0]) && b != rune(one[1]) {
			fourX = append(fourX, b)
		}
	}

	count := 0
	for _, b := range s {
		if b == fourX[0] || b == fourX[1] {
			count++
		}
	}

	return count == 2
}

func isSix(one, s string) bool {
	count := 0

	for _, b := range []byte(s) {
		if b == one[0] || b == one[1] {
			count++
		}
	}

	return count == 1
}

func isNine(four, s string) bool {
	count := 0

	for _, sr := range s {
		for _, fr := range four {
			if sr == fr {
				count++
			}
		}
	}

	return count == len(four)
}

// func createNewNumsMap(n8 string) map[string]int {
// 	m := make(map[string]int)
// 	std8 := numsToLetters[8]

// 	n8map := make(map[string]string)

// 	for i := 0; i < 7; i++ {
// 		n8map[string(std8[i])] = string(n8[i])
// 	}

// 	for i, sNum := range numsToLetters {
// 		var sb strings.Builder

// 		for j := 0; j < len(sNum); j++ {
// 			key := string(sNum[j])
// 			v := n8map[key]
// 			sb.WriteString(v)
// 		}

// 		m[sb.String()] = i
// 	}

// 	return m
// }
