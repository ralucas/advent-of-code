package day17_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	day17 "github.com/ralucas/advent-of-code/pkg/2020/17"
)

//nolint
var td day17.Day

// func TestMain(m *testing.M) {
// 	td.PrepareData("../../../test/testdata/2020/17/test_input.txt")

// 	m.Run()

// 	os.Exit(0)
// }

func TestPrepareData(t *testing.T) {
	t.Skip("not finished")
	assert.Equal(t, 3, len(td.Cubes))
}
