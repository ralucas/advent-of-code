package day13

import "strconv"

type Element struct {
	value    *int
	children []*Element
}

func Parse(s string) ([]*Element, error) {
	output := make([]*Element, 0)

	i := 1
	for i < len(s) {
		b := s[i]

		if b == ClosingBracket {
			return output, nil
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

			output = append(output, &Element{value: &val})
		}

		if b == OpeningBracket {
			node := &Element{}

			children, err := Parse(s[i:])
			if err != nil {
				return nil, err
			}

			node.children = children
			output = append(output, node)

			i += (len(children) * 2) + 1
		}

		i += 1
	}

	return output, nil
}

func (e *Element) Value() int {
	return *e.value
}

func (e *Element) Children() []*Element {
	return e.children
}

func (e *Element) IsLeaf() bool {
	return e.value != nil
}
