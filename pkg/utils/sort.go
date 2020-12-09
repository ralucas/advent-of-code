package utils

func QSort(vi []int) []int {
	vlen := len(vi)

	if vlen <= 1 {
		return vi
	}

	if vlen == 2 {
		if vi[0] > vi[1] {
			return []int{vi[1], vi[0]}
		}
		return vi
	}

	pivot := vlen / 2

	a, b := make([]int, 0), make([]int, 0)

	// Always remember pivot can't be included in
	// recursive array
	for _, n := range vi {
		if n < vi[pivot] {
			a = append(a, n)
		}
		if n > vi[pivot] {
			b = append(b, n)
		}
	}

	sorted := make([]int, 0)

	sorted = append(sorted, QSort(a)...)
	sorted = append(sorted, vi[pivot])
	sorted = append(sorted, QSort(b)...)

	return sorted
}
