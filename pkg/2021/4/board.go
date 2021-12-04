package day4

type Board struct {
	values [][]int
	marks  [][]int
}

func NewBoard(values [][]int) *Board {
	return &Board{
		values: values,
		marks:  make([][]int, 0),
	}
}

func (b *Board) Values() [][]int {
	return b.values
}

func (b *Board) mark(val int) bool {
	for i := range b.values {
		for j := range b.values[i] {
			if b.values[i][j] == val {
				b.marks = append(b.marks, []int{i, j})
				return b.bingo()
			}
		}
	}

	return false
}

func (b *Board) bingo() bool {
	return false
}
