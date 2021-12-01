package day7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var td, td1 Day

func init() {
	td.PrepareData("../../test/testdata/7/test_input.txt")
	td1.PrepareData("../../test/testdata/7/test_input1.txt")
}

func TestPrepareData(t *testing.T) {

	assert.NotNil(t, td.data)

	t.Run("has the correct len", func(t *testing.T) {
		assert.Equal(t, 9, len(td.data))
	})

	t.Run("has the correct items", func(t *testing.T) {
		testMap := map[string]map[string]int{
			"light_red": {
				"bright_white": 1,
				"muted_yellow": 2,
			},
			"dark_orange": {
				"bright_white": 3,
				"muted_yellow": 4,
			},
			"bright_white": {
				"shiny_gold": 1,
			},
			"muted_yellow": {
				"shiny_gold": 2,
				"faded_blue": 9,
			},
			"shiny_gold": {
				"dark_olive":   1,
				"vibrant_plum": 2,
			},
			"dark_olive": {
				"faded_blue":   3,
				"dotted_black": 4,
			},
			"vibrant_plum": {
				"faded_blue":   5,
				"dotted_black": 6,
			},
			"faded_blue":   nil,
			"dotted_black": nil,
		}

		for k := range td.data {
			val, ok := testMap[k]
			assert.True(t, ok, fmt.Sprintf("Failed on %s", k))

			for vk := range val {
				vkv, ok := val[vk]
				assert.True(t, ok, fmt.Sprintf("Failed on %s[%s]", k, vk))
				assert.Equal(t, testMap[k][vk], vkv, fmt.Sprintf("Failed on %s[%s] = %d", k, vk, vkv))
			}
		}
	})
}

func TestCountParents(t *testing.T) {
	count, _ := CountParents("shiny_gold", td.data)

	assert.Equal(t, 4, count)
}

func TestCountContains(t *testing.T) {

	t.Run("correctly counts the contains 1", func(t *testing.T) {
		count, _ := CountContains("shiny_gold", td.data)

		assert.Equal(t, 32, count)
	})

	t.Run("correctly counts the contains 2", func(t *testing.T) {
		count, _ := CountContains("shiny_gold", td1.data)

		assert.Equal(t, 126, count)
	})
}
