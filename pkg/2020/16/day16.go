package day16

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	fields   []Field
	myticket Ticket
	tickets  []Ticket
}

type Ticket struct {
	Vals []int
}

type Field struct {
	Name   string
	Ranges [][]int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFile(filepath)
	sections := strings.Split(data, "\n\n")

	fields := make([]Field, 0)
	myticket := Ticket{}
	tickets := make([]Ticket, 0)

	for i, section := range sections {
		spl := strings.Split(section, "\n")

		switch i {
		case 0:
			for _, s := range spl {
				d := strings.Split(s, ": ")
				sRanges := strings.Split(d[1], " or ")
				var ranges [][]int
				for _, sr := range sRanges {
					sRange := strings.Split(sr, "-")
					r1, err := strconv.Atoi(sRange[0])
					if err != nil {
						log.Fatalf("error in range reading %v", err)
					}
					r2, err := strconv.Atoi(sRange[1])
					if err != nil {
						log.Fatalf("error in range reading %v", err)
					}
					ranges = append(ranges, []int{r1, r2})
				}
				f := Field{
					Name:   d[0],
					Ranges: ranges,
				}
				fields = append(fields, f)
			}
		case 1:
			myticket.Vals = arrayutil.MapToInt(strings.Split(spl[1], ","))
		case 2:
			for _, s := range spl[1:] {
				t := Ticket{arrayutil.MapToInt(strings.Split(s, ","))}
				tickets = append(tickets, t)
			}
		}
	}

	d.fields = fields
	d.myticket = myticket
	d.tickets = tickets

	return
}

func (d *Day) Part1() interface{} {
	ff := BuildFieldFilter(d.fields)
	return SumInvalidTickets(ff, d.tickets)
}

func (d *Day) Part2() interface{} {
	colsmap := make(map[string][]int)
	for _, field := range d.fields {
		colsmap[field.Name] = FindColumnsByFieldName(*d, field.Name)
	}

	colmap := DiscoverColumns(colsmap)

	depRegex := regexp.MustCompile(`^departure`)

	total := 1

	for name, val := range colmap {
		if depRegex.Match([]byte(name)) {
			total *= d.myticket.Vals[val]
		}
	}

	return total
}

func BuildFieldFilter(fields []Field, fieldNames ...string) []int {
	min, max := fields[0].Ranges[0][0], fields[0].Ranges[0][1]

	var filteredFields []Field

	if fieldNames != nil {
		for _, field := range fields {
			if arrayutil.Index(fieldNames, field.Name) != -1 {
				filteredFields = append(filteredFields, field)
			}
		}
	} else {
		filteredFields = fields
	}

	for _, field := range filteredFields {
		for _, r := range field.Ranges {
			if r[0] < min {
				min = r[0]
			}
			if r[1] > max {
				max = r[1]
			}
		}
	}

	filter := make([]int, max+1)

	for _, field := range filteredFields {
		for _, r := range field.Ranges {
			for i := r[0]; i <= r[1]; i++ {
				filter[i] = 1
			}
		}
	}

	return filter
}

func SumInvalidTickets(filter []int, tickets []Ticket) int {
	sum := 0
	flen := len(filter)

	for _, ticket := range tickets {
		for _, n := range ticket.Vals {
			if n >= flen || filter[n] == 0 {
				sum += n
			}
		}
	}

	return sum
}

func FilterInvalidTickets(filter []int, tickets []Ticket) []Ticket {
	flen := len(filter)

	filteredTickets := make([]Ticket, 0)

	for _, ticket := range tickets {
		hasEvery := arrayutil.Every(ticket.Vals, func(v int, _ int) bool {
			return v < flen && filter[v] != 0
		})
		if hasEvery {
			filteredTickets = append(filteredTickets, ticket)
		}
	}

	return filteredTickets
}

func FindColumnsByFieldName(data Day, fieldName string) []int {
	ff := BuildFieldFilter(data.fields, fieldName)
	flen := len(ff)

	idxs := make([]int, len(data.myticket.Vals))

	filteredTickets := FilterInvalidTickets(
		BuildFieldFilter(data.fields), data.tickets,
	)

	for _, ticket := range filteredTickets {
		for i, v := range ticket.Vals {
			if v >= flen || ff[v] == 0 {
				idxs[i] = -1
			}
		}
	}

	return arrayutil.IndexesInt(idxs, 0)
}

func DiscoverColumns(cols map[string][]int) map[string]int {
	clen := len(cols)

	cmap := make(map[string]int)

	for {
		var found []string
		for k, v := range cols {
			if len(v) == 1 {
				found = append(found, k)
			}
		}

		for _, f := range found {
			vi, _ := cols[f]
			foundval := vi[0]
			cmap[f] = foundval
			for k, v := range cols {
				if len(v) > 1 {
					var newvals []int
					for _, vv := range v {
						if vv != foundval {
							newvals = append(newvals, vv)
						}
					}
					cols[k] = newvals
				}
			}
		}

		if clen == len(found) {
			return cmap
		}

	}
}
