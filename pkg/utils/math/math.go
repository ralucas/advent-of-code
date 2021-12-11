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

func Extent(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func Mean(vi []int) int {
	sum := Sum(vi)
	return sum / len(vi)
}

// expects sorted array
func Median(vi []int) int {
	n := len(vi)
	mid := n / 2

	if n%2 == 0 {
		return (vi[mid] + vi[mid+1]) / 2
	}

	return vi[mid]
}
