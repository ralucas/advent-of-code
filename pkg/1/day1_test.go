package day1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ralucas/advent-of-code/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var testDay Day

func init() {
	testDay.PrepareData("../../assets/1/input.txt")
}

func TestTwoSum(t *testing.T) {
	dlen := len(testDay.data)

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
		targ := testDay.data[x] + testDay.data[y]
		a, b := TwoSum(testDay.data, targ)
		assert.Equal(t, targ, a+b, fmt.Sprintf("got %d + %d, expected %d + %d = %d", a, b, x, y, targ))
	}
}

func TestTwoSumSorted(t *testing.T) {
	sData := utils.QSort(testDay.data)
	dlen := len(testDay.data)

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
		targ := testDay.data[x] + testDay.data[y]
		a, b := TwoSumSorted(sData, targ)
		assert.Equal(t, targ, a+b, fmt.Sprintf("got %d + %d, expected %d + %d = %d", a, b, x, y, targ))
	}
}

func TestThreeSum(t *testing.T) {
	sData := utils.QSort(testDay.data)
	dlen := len(testDay.data)
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
		targ := testDay.data[x] + testDay.data[y] + testDay.data[z]
		a, b, c := ThreeSum(sData, targ)
		assert.Equal(t, targ, a+b+c, fmt.Sprintf(
			"got %d, %d, %d, expected: %d + %d + %d = %d",
			a, b, c, testDay.data[x], testDay.data[y], testDay.data[z], targ))
	}
}
