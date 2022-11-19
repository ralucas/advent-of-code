package day5

import (
	"fmt"
	"log"
	"strings"
	"sync"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type BoardingPass struct {
	Row int
	Col int
	ID  int
}

type Plane struct {
	seating [][]int
	rows    []int
}

type Day struct {
	data [][]string
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	prepared := make([][]string, len(data))

	for i, d := range data {
		prepared[i] = strings.Split(d, "")
	}

	d.data = prepared

	return
}

func (d *Day) Part1() interface{} {
	maxID := GetHighestSeatID(d.data)

	return maxID
}

func (d *Day) Part2() interface{} {
	mySeat := FindSeat(d.data)

	return mySeat.ID
}

func getRow(s []string) int {
	return getNum(0, 127, "F", "B", s)
}

func getCol(s []string) int {
	return getNum(0, 7, "L", "R", s)
}

func getNum(min, max int, minL, maxL string, s []string) int {
	for _, c := range s[:len(s)-1] {
		if c != minL && c != maxL {
			errMsg := fmt.Errorf("Bad string pass sent in, got %s", c)
			log.Println(errMsg)
			return -1
		}
		switch c {
		case minL:
			max = (min + max) / 2
		case maxL:
			min = ((min + max) / 2) + 1
		}
	}

	if s[len(s)-1] == minL {
		return min
	}

	return max
}

// toBoardingPass converts a boarding pass string to numeric values.
// Every seat also has a unique seat ID: multiply the row by 8, then add the column
func toBoardingPass(s []string) BoardingPass {
	row := getRow(s[:7])
	col := getCol(s[7:])

	bp := BoardingPass{
		Row: row,
		Col: col,
		ID:  (row * 8) + col,
	}

	return bp
}

func GetHighestSeatID(ss [][]string) int {
	var maxID int

	for _, s := range ss {
		bp := toBoardingPass(s)
		if bp.ID > maxID {
			maxID = bp.ID
		}
	}

	return maxID
}

func NewPlane(ss [][]string) Plane {
	plane := Plane{
		seating: make([][]int, 128),
		rows:    make([]int, 128),
	}
	for i := range plane.seating {
		plane.seating[i] = make([]int, 8)
	}

	// Setup wait group for par-for
	wg := sync.WaitGroup{}

	wg.Add(len(ss))

	for _, s := range ss {
		go func(spass []string) {
			bp := toBoardingPass(spass)
			plane.seating[bp.Row][bp.Col] = 1
			plane.rows[bp.Row] += 1
			wg.Done()
		}(s)
	}

	wg.Wait()

	return plane
}

func (p *Plane) Print() {
	for i, row := range p.seating {
		fmt.Println(i, ":", row)
	}
}

func findAvailableSeats(plane Plane) []BoardingPass {
	possibleSeats := []BoardingPass{}

	for i, row := range plane.seating {
		if plane.rows[i] < 8 {
			availableSeats := arrayutil.FindIntIndexes(row, func(v int) bool { return v == 0 })
			if len(availableSeats) > 0 {
				for _, as := range availableSeats {
					bp := BoardingPass{
						Row: i,
						Col: as,
						ID:  (i * 8) + as,
					}
					possibleSeats = append(possibleSeats, bp)
				}
			}
		}
	}

	return possibleSeats
}

func FindSeat(ss [][]string) BoardingPass {
	plane := NewPlane(ss)

	availSeats := findAvailableSeats(plane)

	var mySeat BoardingPass

	for _, seat := range availSeats {
		if seat.Row != 0 && seat.Row != 127 {
			// After first match, let's call it the seat
			return seat
		}
	}

	return mySeat
}
