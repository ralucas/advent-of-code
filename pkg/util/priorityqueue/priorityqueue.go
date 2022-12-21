package priorityqueue

import (
	"fmt"
)

type PriorityQueueKind int

const (
	MaxPriorityQueue PriorityQueueKind = iota
	MinPriorityQueue
)

type Item interface {
	String() string
}

type value struct {
	item   Item
	weight int
}

type PriorityQueue struct {
	heap []value
	kind PriorityQueueKind
}

func New(k PriorityQueueKind) *PriorityQueue {
	return &PriorityQueue{
		heap: make([]value, 0),
		kind: k,
	}
}

func NewMaxPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: make([]value, 0),
		kind: MaxPriorityQueue,
	}
}

func NewMinPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: make([]value, 0),
		kind: MinPriorityQueue,
	}
}

func (pq *PriorityQueue) String() []string {
	sHeap := make([]string, 0)

	for _, val := range pq.heap {
		s := fmt.Sprintf("%s=%d", val.item.String(), val.weight)
		sHeap = append(sHeap, s)
	}

	return sHeap
}

// Len returns the length of the queue
func (pq *PriorityQueue) Len() int {
	return len(pq.heap)
}

// Empty returns whether or not the queue is empty
func (pq *PriorityQueue) Empty() bool {
	return len(pq.heap) == 0
}

// weight returns the weight given an index
func (pq *PriorityQueue) weight(index int) int {
	val := pq.heap[index]

	return val.weight
}

// parentIndex returns the parent index given an index
func (pq *PriorityQueue) parentIndex(index int) int {
	return (index - 1) / 2
}

func (pq *PriorityQueue) leftChildIndex(index int) int {
	return (2 * index) + 1
}

func (pq *PriorityQueue) rightChildIndex(index int) int {
	return (2 * index) + 2
}

func (pq *PriorityQueue) siftDownMax(index int) {
	maxIdx := index

	lc := pq.leftChildIndex(index)
	if lc < len(pq.heap) && (pq.weight(lc) > pq.weight(maxIdx)) {
		maxIdx = lc
	}

	rc := pq.rightChildIndex(index)
	if rc < len(pq.heap) && (pq.weight(rc) > pq.weight(maxIdx)) {
		maxIdx = rc
	}

	if index != maxIdx {
		pq.heap[maxIdx], pq.heap[index] = pq.heap[index], pq.heap[maxIdx]
		pq.siftDownMax(maxIdx)
	}
}

func (pq *PriorityQueue) siftDownMin(index int) {
	minIdx := index

	lc := pq.leftChildIndex(index)
	if lc < len(pq.heap) && (pq.weight(lc) < pq.weight(minIdx)) {
		minIdx = lc
	}

	rc := pq.rightChildIndex(index)
	if rc < len(pq.heap) && (pq.weight(rc) < pq.weight(minIdx)) {
		minIdx = rc
	}

	if index != minIdx {
		pq.heap[minIdx], pq.heap[index] = pq.heap[index], pq.heap[minIdx]
		pq.siftDownMin(minIdx)
	}
}

func (pq *PriorityQueue) siftUpMax(index int) {
	child := index
	// sift up
	for {
		pi := pq.parentIndex(child)
		if pq.weight(pi) < pq.weight(child) {
			// swap
			pq.heap[pi], pq.heap[child] = pq.heap[child], pq.heap[pi]

			child = pi
		} else {
			return
		}
	}
}

func (pq *PriorityQueue) siftUpMin(index int) {
	child := index
	// sift up
	for child > 0 {
		pi := pq.parentIndex(child)
		if pq.weight(pi) > pq.weight(child) {
			// swap
			pq.heap[pi], pq.heap[child] = pq.heap[child], pq.heap[pi]

			child = pi
		} else {
			return
		}
	}
}

func (pq *PriorityQueue) heapify(count int) {
	end := 1

	for end < count {
		if pq.kind == MaxPriorityQueue {
			pq.siftUpMax(end)
		} else {
			pq.siftUpMin(end)
		}
		end += 1
	}
}

func (pq *PriorityQueue) Pop() (Item, int) {
	if len(pq.heap) == 0 {
		return nil, 0
	}

	val := pq.heap[0]

	// put the end at the front, shrink the heap, then sift down
	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]

	pq.heapify(len(pq.heap))

	return val.item, val.weight
}

func (pq *PriorityQueue) Insert(node Item, weight int) {
	v := value{node, weight}

	pq.heap = append(pq.heap, v)

	pq.heapify(len(pq.heap))
}
