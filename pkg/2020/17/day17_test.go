package day17_test

import (
	"os"
	"testing"

	day17 "github.com/ralucas/advent-of-code/pkg/17"
	"github.com/stretchr/testify/assert"
)

//nolint
var td day17.Day

func TestMain(m *testing.M) {
	td.PrepareData("../../test/testdata/17/test_input.txt")

	m.Run()

	os.Exit(0)
}

func TestPrepareData(t *testing.T) {
	assert.Equal(t, 3, len(td.Cubes))
}
