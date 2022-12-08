package day8

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	"github.com/ralucas/advent-of-code/pkg/util/queue"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	mathutil "github.com/ralucas/advent-of-code/pkg/util/math"
)

type Day struct {
	grid  [][]int
	gridT [][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")
	data1 := arrayutil.MapTo2D(data, "")

	d.grid = make([][]int, len(data1))

	for i := range data1 {
		d.grid[i] = arrayutil.MapToInt(data1[i])
	}

	d.gridT = arrayutil.Transpose(d.grid)

	return
}

func (d *Day) IsVisible(height, row, col int) bool {
	maxLeft := mathutil.Max(d.grid[row][:col]...)
	maxRight := mathutil.Max(d.grid[row][col+1:]...)
	maxUp := mathutil.Max(d.gridT[col][:row]...)
	maxDown := mathutil.Max(d.gridT[col][row+1:]...)

	return height > maxLeft || height > maxRight || height > maxUp || height > maxDown
}

func (d *Day) NextForQueue(row, col int, seen [][]int) (rowCols [][]int) {
	maxRow := len(d.grid) - 1
	maxCol := len(d.grid[0]) - 1

	next := make([][]int, 0)
	if row+1 < maxRow && seen[row+1][col] == 0 {
		next = append(next, []int{d.grid[row+1][col], row + 1, col})
	}

	if col+1 < maxCol && seen[row][col+1] == 0 {
		next = append(next, []int{d.grid[row][col+1], row, col + 1})
	}

	return next
}

func (d *Day) createVisibleGrid() (visible [][]int, seen [][]int, initialCount int) {
	// set at -4 because the 4 corners are counted twice in iteration
	visibleCount := -4

	visibleGrid := make([][]int, len(d.grid))
	seenGrid := make([][]int, len(d.grid))
	for i := range d.grid {
		visibleGrid[i] = make([]int, len(d.grid[i]))
		seenGrid[i] = make([]int, len(d.grid[i]))
		visibleGrid[i][0] = 1
		seenGrid[i][0] = 1
		visibleGrid[i][len(d.grid)-1] = 1
		seenGrid[i][len(d.grid)-1] = 1
		visibleCount += 2
		if i == 0 || i == len(d.grid)-1 {
			for j := range d.grid[i] {
				visibleGrid[i][j] = 1
				seenGrid[i][j] = 1
				visibleCount += 1
			}
		}
	}

	return visibleGrid, seenGrid, visibleCount
}

func (d *Day) Part1() interface{} {
	visibleGrid, seenGrid, visibleCount := d.createVisibleGrid()

	q := queue.New([]int{d.grid[1][1], 1, 1})

	for !q.Empty() {
		item, err := q.Pop()
		if err != nil {
			log.Fatal(err)
		}

		height, row, col := item[0], item[1], item[2]

		if seenGrid[row][col] == 1 {
			continue
		}

		seenGrid[row][col] = 1

		if d.IsVisible(height, row, col) {
			visibleGrid[row][col] = 1
			visibleCount += 1
		}

		next := d.NextForQueue(row, col, seenGrid)
		if len(next) > 0 {
			q.Push(next...)
		}
	}

	return visibleCount
}

func (d *Day) CountGreaterThan(val int, arr []int, dir int) int {
	count := 0

	if dir == 1 {
		for i := range arr {
			if val <= arr[i] {
				return count + 1
			}
			count += 1
		}
	}

	if dir == -1 {
		for i := len(arr) - 1; i >= 0; i-- {
			if val <= arr[i] {
				return count + 1
			}
			count += 1
		}
	}

	return count
}

func (d *Day) ScenicScore(height, row, col int) int {
	countLeft := d.CountGreaterThan(height, d.grid[row][:col], -1)
	countRight := d.CountGreaterThan(height, d.grid[row][col+1:], 1)
	countUp := d.CountGreaterThan(height, d.gridT[col][:row], -1)
	countDown := d.CountGreaterThan(height, d.gridT[col][row+1:], 1)

	return countLeft * countRight * countUp * countDown
}

func (d *Day) Part2() interface{} {
	_, seenGrid, _ := d.createVisibleGrid()

	maxScenicScore := 0

	q := queue.New([]int{d.grid[1][1], 1, 1})

	for !q.Empty() {
		item, err := q.Pop()
		if err != nil {
			log.Fatal(err)
		}

		height, row, col := item[0], item[1], item[2]

		if seenGrid[row][col] == 1 {
			continue
		}

		seenGrid[row][col] = 1

		score := d.ScenicScore(height, row, col)
		if score > maxScenicScore {
			maxScenicScore = score
		}

		next := d.NextForQueue(row, col, seenGrid)
		if len(next) > 0 {
			q.Push(next...)
		}
	}

	return maxScenicScore
}
