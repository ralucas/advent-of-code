package math

func Sum(vi []int) int {
	sum := 0
	for _, v := range vi {
		sum += v
	}

	return sum
}
