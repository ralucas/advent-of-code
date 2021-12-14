package day12_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	day12 "github.com/ralucas/advent-of-code/pkg/2021/12"
)

var td1 day12.Day
var td2 day12.Day
var td3 day12.Day

func TestMain(m *testing.M) {
	m.Run()
}

func TestPart1(t *testing.T) {
	t.Run("CorrectNumberOfPaths", func(t *testing.T) {
		t.Parallel()

		days := []day12.Day{td1, td2, td3}
		expects := []int{10, 19, 226}

		for i, day := range days {
			t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
				day.PrepareData(fmt.Sprintf("../../../test/testdata/2021/12/test_input%d.txt", i+1))
				result := day.Part1()
				assert.Equal(t, expects[i], result)
			})
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("CorrectNumberOfPaths", func(t *testing.T) {
		t.Parallel()

		days := []day12.Day{
			td1,
			td2,
			td3,
		}
		expects := []int{
			36,
			103,
			3509,
		}

		for i, day := range days {
			t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
				day.PrepareData(fmt.Sprintf("../../../test/testdata/2021/12/test_input%d.txt", i+1))
				result := day.Part2()
				assert.Equal(t, expects[i], result)
			})
		}
	})
}
