package day16_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	day16 "github.com/ralucas/advent-of-code/pkg/2021/16"
)

var td day16.Day

func TestMain(m *testing.M) {
	m.Run()
}

func TestPart1(t *testing.T) {
	td.SetData(strings.Split("8A004A801A8002F478", ""))
	result := td.Part1()
	expect := 16

	assert.Equal(t, expect, result)
}

func TestPart2(t *testing.T) {
	result := td.Part2()
	expect := true

	assert.Equal(t, expect, result)
}
