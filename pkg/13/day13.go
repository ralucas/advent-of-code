package day13

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	earliestTimestamp int
	busSchedule       []string
	bsMap             []map[int]int
	buses             []int
}

const MaxUint64 uint64 = 1<<64 - 1
const MinInt64 int64 = (1<<48 - 1)
const MaxInt64 int64 = (1<<63 - 1)

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	var err error
	d.earliestTimestamp, err = strconv.Atoi(data[0])
	if err != nil {
		log.Fatalf("Failed to parse earliest timestamp")
	}

	busSchedule := strings.Split(data[1], ",")
	d.busSchedule = busSchedule

	d.bsMap = createScheduleMap(busSchedule)

	filtered := utils.Filter(busSchedule, func(s string) bool {
		return s != "x"
	})

	d.buses = utils.MapToInt(filtered)

	return
}

func createScheduleMap(vs []string) []map[int]int {
	output := make([]map[int]int, 0)
	count := 0
	for _, b := range vs {
		if b == "x" {
			count++
		} else {
			n, err := strconv.Atoi(b)
			if err != nil {
				log.Fatalf("atoi error %v", err)
			}
			m := make(map[int]int)
			m[n] = count + 1
			output = append(output, m)
			count = 0
		}
	}

	return output
}

func (d *Day) Part1() interface{} {
	min, minBus := NearestNextBus(d.earliestTimestamp, d.buses)

	return min * minBus
}

func (d *Day) Part2() interface{} {
	val := EarliestTimestampForScheduleSlow(d.bsMap)

	return val
}

func NearestNextBus(ts int, buses []int) (int, int) {
	min := int(^uint(0) >> 1)
	minBus := buses[0]

	for _, bus := range buses {
		div := ts / bus
		next := (div + 1) * bus
		diff := next - ts
		if diff < min {
			min = diff
			minBus = bus
		}
	}

	return min, minBus
}

func keyValue(m map[int]int) (int, int) {
	for k, v := range m {
		return k, v
	}

	return -1, -1
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	mid := n / 2

	for i := 3; i < mid; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primeFactors(val int) []int {
	if val <= 1 {
		return []int{}
	}

	if val <= 3 {
		return []int{val}
	}

	mid := val / 2

	factors := make([]int, 0)

	if val%2 == 0 {
		factors = append(factors, 2)
		factors = append(factors, primeFactors(val/2)...)
		return factors
	}

	for i := 3; i < mid; i += 2 {
		if val%i == 0 {
			if isPrime(i) {
				factors = append(factors, i)
				factors = append(factors, primeFactors(val/i)...)
				break
			} else {
				factors = append(factors, primeFactors(i)...)
			}
			if mid > val/i {
				mid = (val / i) - 1
			}
		}
	}

	if isPrime(val) {
		factors = append(factors, val)
	}

	return factors
}

func leastCommonMultiples(vi []int) int {
	pfMap := make(map[int]int)
	for _, v := range vi {
		pfs := primeFactors(v)
		cpfMap := make(map[int]int)
		for _, pf := range pfs {
			cpfMap[pf] += 1
		}
		for key, val := range cpfMap {
			if pval, ok := pfMap[key]; ok {
				if val > pval {
					pfMap[key] = val
				}
			} else {
				pfMap[key] = val
			}
		}
	}

	lcm := 1
	for key, val := range pfMap {
		cur := 1
		for i := 0; i < val; i++ {
			cur *= key
		}
		lcm *= cur
	}

	return lcm
}

func EarliestTimestampForSchedule(schedule []map[int]int) int {
	var scheds []int
	bus, _ := keyValue(schedule[0])
	rems := []int{bus}
	totalSkip := 0
	for i := 1; i < len(schedule); i++ {
		nbus, skip := keyValue(schedule[i])
		totalSkip += skip
		rems = append(rems, totalSkip)
		scheds = append(scheds, nbus)
	}

	return leastCommonMultiples(rems)
}

func EarliestTimestampForScheduleSlow(schedule []map[int]int) int64 {
	var rems []int64
	var scheds []int64
	bus, _ := keyValue(schedule[0])
	firstBus := int64(bus)
	totalSkip := 0
	for i := 1; i < len(schedule); i++ {
		nbus, skip := keyValue(schedule[i])
		totalSkip += skip
		rems = append(rems, int64(nbus-totalSkip))
		scheds = append(scheds, int64(nbus))
	}

	maxProcs := int64(runtime.GOMAXPROCS(0))
	starts := make([]int64, maxProcs)
	starts[0] = MinInt64
	sect := (MaxInt64 - starts[0]) / maxProcs
	for i := 1; i < int(maxProcs)-1; i++ {
		starts[i] = starts[i-1] + sect
	}
	starts[len(starts)-1] = MaxInt64

	done := make(chan bool, len(starts))
	result := make(chan int64, len(starts))
	fmt.Printf("%d :: %d %d\n", firstBus, starts[0], starts[len(starts)-1])

	for i := 0; i < len(starts)-1; i++ {
		fmt.Printf("%d %d :: %d %d\n", i, firstBus, starts[i], starts[i+1])

		go findLowestCommonNumerator(
			starts[i],
			starts[i+1],
			firstBus,
			result,
			done,
			rems,
			scheds)
	}

	output := MaxInt64
	for i := 0; i < len(starts)-1; i++ {
		d := <-done
		if d {
			r := <-result
			fmt.Printf("result for %d : %d\n", i, r)
			if r != int64(0) && r < output {
				output = r
			}
		} else {
			fmt.Printf("%d :: %d\n", <-result, i)
		}
	}

	if output == MaxInt64 {
		return int64(-1)
	}

	return output
}

func findLowestCommonNumerator(
	start int64,
	end int64,
	bus int64,
	resultCh chan int64,
	doneCh chan bool,
	remainders []int64,
	buses []int64,
) {

	x := start / bus
	if x < 1 {
		x = 1
	}
	y := end / bus

	rs := x * bus

	rlen := len(remainders)
	fmt.Printf("%d :: %d %d\n", bus, x, y)

	progress := float64(0)

	for x < y {

		progress = float64((x - start) / y)

		rs = x * bus
		for i, rem := range remainders {
			if rs < buses[i] {
				break
			}
			if rs%buses[i] != rem {
				break
			}
			if i == rlen-1 {
				doneCh <- true
				resultCh <- rs
			}
		}
		x++
	}
	doneCh <- true
	resultCh <- 0

	return
}
