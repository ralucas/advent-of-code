package day1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	sorting "github.com/ralucas/advent-of-code/pkg/util/sort"
)

var td Day

func init() {
	td.PrepareData("../../../assets/2020/1/input.txt")
}

func TestTwoSum(t *testing.T) {
	dlen := len(td.data)

	x, y := 0, 0

	for i := 0; i < 100000; i++ {
		for {
			time.Sleep(time.Nanosecond * 100)
			rand.Seed(time.Now().UnixNano())
			x, y = rand.Intn(dlen), rand.Intn(dlen)
			if x != y {
				break
			}
		}
		targ := td.data[x] + td.data[y]
		a, b := TwoSum(td.data, targ)
		assert.Equal(t, targ, a+b, fmt.Sprintf("got %d + %d, expected %d + %d = %d", a, b, x, y, targ))
	}
}

func TestTwoSumSorted(t *testing.T) {
	sData := sorting.QSort(td.data)
	dlen := len(td.data)

	x, y := 0, 0

	for i := 0; i < 100000; i++ {
		for {
			time.Sleep(time.Nanosecond * 100)
			rand.Seed(time.Now().UnixNano())
			x, y = rand.Intn(dlen), rand.Intn(dlen)
			if x != y {
				break
			}
		}
		targ := td.data[x] + td.data[y]
		a, b := TwoSumSorted(sData, targ)
		assert.Equal(t, targ, a+b, fmt.Sprintf("got %d + %d, expected %d + %d = %d", a, b, x, y, targ))
	}
}

func TestThreeSum(t *testing.T) {
	sData := sorting.QSort(td.data)
	dlen := len(td.data)
	x, y, z := 0, 0, 0

	for i := 0; i < 100000; i++ {
		for {
			time.Sleep(time.Nanosecond)
			rand.Seed(time.Now().UnixNano())
			x, y, z = rand.Intn(dlen), rand.Intn(dlen), rand.Intn(dlen)
			if x != y && y != z && x != z {
				break
			}
		}
		targ := td.data[x] + td.data[y] + td.data[z]
		a, b, c := ThreeSum(sData, targ)
		assert.Equal(t, targ, a+b+c, fmt.Sprintf(
			"got %d, %d, %d, expected: %d + %d + %d = %d",
			a, b, c, td.data[x], td.data[y], td.data[z], targ))
	}
}
