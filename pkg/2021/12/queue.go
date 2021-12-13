package day12

type Queue struct {
	values []string
}

func NewQueue(values ...string) *Queue {
	return &Queue{
		values: values,
	}
}

func (q *Queue) Pop() string {
	val := q.values[0]
	q.values = q.values[1:]

	return val
}

func (q *Queue) Push(val string) {
	q.values = append(q.values, val)
}

func (q *Queue) Empty() bool {
	return len(q.values) == 0
}
