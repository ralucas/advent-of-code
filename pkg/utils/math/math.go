package math

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Sum(vi []int) int {
	sum := 0
	for _, v := range vi {
		sum += v
	}

	return sum
}

func Max(vi ...int) int {
	max := MinInt
	for _, v := range vi {
		if v > max {
			max = v
		}
	}

	return max
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}