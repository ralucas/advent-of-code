package day11

import (
	"log"
	"sync"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

const (
	EmptySeat    = "L"
	Floor        = "."
	OccupiedSeat = "#"
)

type Day struct {
	data [][]string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	out := utils.MapTo2D(data, "")

	d.data = out

	return
}

func (d *Day) Part1() interface{} {
	input := d.data
	filled := Fill(input)
	for {
		if IsEqual(filled, input) {
			break
		}
		input = filled
		filled = Fill(input)
	}

	occupied, _ := CountSeats(filled)
	return occupied
}

func (d *Day) Part2() interface{} {
	return -1
}

// "Adjacent" to a given seat means one of the eight positions immediately
// up, down, left, right, or diagonal from the seat.
//
// The following rules are applied to every seat simultaneously:
// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
// Otherwise, the seat's state does not change.

// GetSeatState checks adjacency of a seat
// given the i, j of the seat according to the
// rules above.
// returns OccupiedSeat if no occupied seats adjacent
// returns EmptySeat if seat is occupied and four or more seats adjacent to it are also occupied
// returns the same input if neither
func GetSeatState(i, j int, vss [][]string) string {
	adj := [][]int{
		{i - 1, j},     //up
		{i + 1, j},     //down
		{i, j - 1},     //left
		{i, j + 1},     //right
		{i - 1, j - 1}, //up-left
		{i - 1, j + 1}, //up-right
		{i + 1, j - 1}, //down-left
		{i + 1, j + 1}, //down-right
	}

	vlen := len(vss)
	svlen := len(vss[0])

	occCount := 0
	empCount := 0
	for _, ij := range adj {
		x := ij[0]
		y := ij[1]

		if x < 0 || x > vlen-1 || y < 0 || y > svlen-1 {
			continue
		}

		switch vss[x][y] {
		case OccupiedSeat:
			occCount += 1
		case EmptySeat:
			empCount += 1
		default:
			continue
		}
	}

	switch vss[i][j] {
	case EmptySeat:
		if occCount == 0 {
			return OccupiedSeat
		}
		return EmptySeat
	case OccupiedSeat:
		if occCount >= 4 {
			return EmptySeat
		}
		return OccupiedSeat
	case Floor:
		return Floor
	}

	// Should never be reached.
	return vss[i][j]
}

// Fill takes a seat pattern and in parallel
// fills the seats given the rules above
// returning a filled seat pattern.
func Fill(vss [][]string) [][]string {
	// Need to properly copy a 2D array
	// via looping and copying.
	// Straight up copy does not work on 2D.
	cpvss := make([][]string, len(vss))
	for i := range cpvss {
		cp := make([]string, len(vss[i]))
		copy(cp, vss[i])
		cpvss[i] = cp
	}

	wg := sync.WaitGroup{}
	for i := range cpvss {
		wg.Add(1)
		go func(idx int) {
			for j := range cpvss[idx] {
				ss := GetSeatState(idx, j, vss)
				cpvss[idx][j] = ss
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return cpvss
}

// IsEqual compares 2 2D arrays to check for equality.
func IsEqual(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

// CountSeats takes in a seat pattern
// and returns OccupiedSeat count, EmptySeat count
func CountSeats(vss [][]string) (int, int) {
	occCount, empCount := 0, 0

	ch := make(chan []int)

	for i := range vss {
		go func(idx int) {
			occ := 0
			emp := 0

			for _, s := range vss[idx] {
				switch s {
				case OccupiedSeat:
					occ += 1
				case EmptySeat:
					emp += 1
				default:
					continue
				}
			}
			ch <- []int{occ, emp}
		}(i)
	}

	for range vss {
		counts := <-ch
		occCount += counts[0]
		empCount += counts[1]
	}

	defer close(ch)
	return occCount, empCount
}
