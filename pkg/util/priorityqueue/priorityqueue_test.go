package priorityqueue_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ralucas/advent-of-code/pkg/util/priorityqueue"
)

type testItem struct {
	x int
	y int
}

func (i *testItem) String() string {
	return fmt.Sprintf("(%d,%d)", i.x, i.y)
}

func TestPriorityQueue(t *testing.T) {
	t.Run("MaxPriorityQueue", func(t *testing.T) {
		pq := priorityqueue.New(priorityqueue.MaxPriorityQueue)

		weights := []int{7, 31, 45, 14, 13, 7, 11, 20, 12}
		for i := range weights {
			pq.Insert(&testItem{i, i}, weights[i])
		}

		assert.Equal(t, 9, pq.Len())

		expects := []int{45, 31, 20, 14, 13, 12, 11, 7, 7}

		for _, expect := range expects {
			node, wt := pq.Pop()
			assert.Equal(t, expect, wt)
			assert.NotNil(t, node)
		}
	})

	t.Run("MinPriorityQueue", func(t *testing.T) {
		pq := priorityqueue.New(priorityqueue.MinPriorityQueue)

		weights := []int{7, 31, 45, 14, 13, 7, 11, 20, 12}
		for i := range weights {
			pq.Insert(&testItem{i, i}, weights[i])
		}

		assert.Equal(t, 9, pq.Len())

		expects := []int{7, 7, 11, 12, 13, 14, 20, 31, 45}

		for _, expect := range expects {
			node, wt := pq.Pop()
			assert.Equal(t, expect, wt)
			assert.NotNil(t, node)
		}
	})

	// t.Run("MinPriorityQueue2", func(t *testing.T) {
	// 	pq := priorityqueue.NewPriorityQueue(priorityqueue.MinPriorityQueue)

	// 	weights := []int{7, 31, 45, 14, 13, 7, 11, 20, 12}
	// 	expects := []int{7, 7, 11, 12, 13, 14, 20, 31, 45}

	// 	for _, wt := range weights {
	// 		pq.Insert(da, weight int)
	// 		assert.Equal(t, expect, wt)
	// 		assert.NotNil(t, node)
	// 	}
	// })
}
