package day13

import "strconv"

const (
	OpeningBracket = byte('[')
	ClosingBracket = byte(']')
	Comma          = byte(',')
)

// List is a doubly linked list, structured as `parent <-> child/parent <-> child`
// such as [1, 2, [3, 4, [5, 6]]], where it's a{1,2} <-> b{3, 4} <-> c{5, 6}
type List struct {
	values   []int
	children []*List
	parent   *List
	index    int
}

func NewList(values ...int) *List {
	return &List{
		values: values,
	}
}

func NewListFromString(s string) (*List, error) {
	root := &List{}
	cur := root

	idx := 0

	i := 1
	for i < len(s) {
		b := s[i]
		if b == OpeningBracket {
			cl := &List{}
			cur.children = append(cur.children, cl)
			cl.parent = cur
			cur = cl
			idx += 1
			cur.index = idx
		}

		if isInt(b) {
			bs := []byte{b}

			for isInt(s[i+1]) {
				bs = append(bs, s[i+1])
				i += 1
			}

			val, err := strconv.Atoi(string(bs))
			if err != nil {
				return nil, err
			}

			cur.values = append(cur.values, val)
		}

		if b == ClosingBracket {
			cur = cur.parent
		}

		i += 1
	}

	return root, nil
}

func isInt(b byte) bool {
	return b >= byte('0') && b <= byte('9')
}

func (l *List) Values() []int {
	return l.values
}

