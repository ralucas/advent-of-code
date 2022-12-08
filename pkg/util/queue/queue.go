package queue

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Any interface {
	string | rune | byte | int | float32 | float64 | []int | []string
}

type Queue[T Any] struct {
	values []T
	ptr    int
	size   int
}

// New creates a new queue with the given values.
func New[T Any](ss ...T) *Queue[T] {
	return &Queue[T]{
		values: ss,
		size:   len(ss),
	}
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.size == 0
}

// Size returns the queue size.
func (q *Queue[T]) Size() int {
	return q.size
}

// Pop removes and returns the top element of the queue.
func (q *Queue[T]) Pop() (T, error) {
	var t T
	if q.Empty() {
		return t, errors.New("empty queue")
	}

	v := q.values[0]

	q.values = q.values[1:]

	q.size -= 1

	return v, nil
}

// PopN removes n elements from the queue.
func (q *Queue[T]) PopN(n int) ([]T, error) {
	if q.Empty() {
		return nil, errors.New("empty queue")
	}

	if q.size < n {
		return nil, errors.New(fmt.Sprintf("queue size [%d] less than requested [%d]", q.size, n))
	}

	vals := q.values[:n]

	q.values = q.values[n:]

	q.size -= n

	return vals, nil
}

// Push adds values to the queue.
func (q *Queue[T]) Push(vals ...T) {
	q.values = append(q.values, vals...)

	q.size += len(vals)
}

// Peek returns the top element of the queue.
func (q *Queue[T]) Peek() T {
	var t T
	if q.Empty() {
		return t
	}

	return q.values[0]
}

// Has returns whether or not the element exists in the queue
// and its first found index
func (q *Queue[T]) Has(element T) (exists bool, index int) {
	for i, v := range q.values {
		if reflect.DeepEqual(v, element) {
			return true, i
		}
	}

	return false, -1
}

// Print returns a string representation of the queue
func (q *Queue[T]) Print() string {
	var sb strings.Builder

	sb.WriteString("[")
	for _, v := range q.values {
		ws := fmt.Sprintf(" %+v", v)
		sb.WriteString(ws)
	}
	sb.WriteString(" ]")

	return sb.String()
}
