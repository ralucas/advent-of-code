package day4

import (
	arrayutil "github.com/ralucas/advent-of-code/pkg/utils/array"
)

type Board struct {
	values [][]int
	marks  [][]int
}

func NewBoard(values [][]int) *Board {
	b := &Board{
		values: values,
	}

	b.marks = make([][]int, len(values))

	for i := range values {
		b.marks[i] = make([]int, len(values[i]))
	}

	return b
}

func (b *Board) Values() [][]int {
	return b.values
}

func (b *Board) UnmarkedValues() []int {
	unmarkedVals := make([]int, 0)
	for i := range b.marks {
		for j := range b.marks[i] {
			if b.marks[i][j] == 0 {
				unmarkedVals = append(unmarkedVals, b.values[i][j])
			}
		}
	}

	return unmarkedVals
}

func (b *Board) Mark(val int) bool {
	for i := range b.values {
		for j := range b.values[i] {
			if b.values[i][j] == val {
				b.marks[i][j] = 1
				return b.bingo()
			}
		}
	}

	return false
}

func checkRowsForBingo(vi [][]int) bool {
	for i := range vi {
		r := arrayutil.Every(vi[i], func(v int, j int) bool {
			return v == 1
		})

		if r {
			return true
		}
	}

	return false
}

func (b *Board) bingo() bool {
	transposedMarks := arrayutil.Transpose(b.marks)

	return checkRowsForBingo(b.marks) ||
		checkRowsForBingo(transposedMarks)
}
