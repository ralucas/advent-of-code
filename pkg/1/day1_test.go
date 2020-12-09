package day1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	util "github.com/ralucas/advent-of-code/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	data := PrepareData("../../assets/1/input.txt")
	dlen := len(data)

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
		targ := data[x] + data[y]
		a, b := TwoSum(data, targ)
		assert.Equal(t, targ, a+b, fmt.Sprintf("got %d + %d, expected %d + %d = %d", a, b, x, y, targ))
	}
}

func TestTwoSumSorted(t *testing.T) {
	data := PrepareData("../../assets/1/input.txt")
	sData := util.QSort(data)
	dlen := len(data)

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
		targ := data[x] + data[y]
		a, b := TwoSumSorted(sData, targ)
		assert.Equal(t, targ, a+b, fmt.Sprintf("got %d + %d, expected %d + %d = %d", a, b, x, y, targ))
	}
}

func TestThreeSum(t *testing.T) {
	data := PrepareData("../../assets/1/input.txt")
	sData := util.QSort(data)
	dlen := len(data)
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
		targ := data[x] + data[y] + data[z]
		a, b, c := ThreeSum(sData, targ)
		assert.Equal(t, targ, a+b+c, fmt.Sprintf(
			"got %d, %d, %d, expected: %d + %d + %d = %d",
			a, b, c, data[x], data[y], data[z], targ))
	}
}
