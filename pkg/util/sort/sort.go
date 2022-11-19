package sort

func QSort(vi []int) []int {
	vlen := len(vi)

	if vlen <= 1 {
		return vi
	}

	if vlen == 2 {
		if vi[0] > vi[1] {
			vi[1], vi[0] = vi[0], vi[1]
		}
		return vi
	}

	qsort(vi, 0, vlen-1)

	return vi
}

// Inplace running of quicksort
func qsort(vi []int, lower int, upper int) {
	if lower >= upper {
		return
	}

	pivot := upper

	i := lower - 1

	for j := lower; j < upper; j++ {
		if vi[j] < vi[pivot] {
			i += 1
			vi[i], vi[j] = vi[j], vi[i]
		}
	}

	i += 1
	vi[i], vi[upper] = vi[upper], vi[i]

	qsort(vi, lower, i-1)
	qsort(vi, i+1, upper)
}

func MergeSort(vi []int) []int {
	vlen := len(vi)

	if vlen <= 1 {
		return vi
	}

	if vlen == 2 {
		if vi[0] > vi[1] {
			vi[1], vi[0] = vi[0], vi[1]
		}
		return vi
	}

	h := vlen / 2

	return merge(MergeSort(vi[:h]), MergeSort(vi[h:]))
}

func merge(via []int, vib []int) []int {
	if len(via) == 0 {
		return vib
	}

	if len(vib) == 0 {
		return via
	}

	if via[0] <= vib[0] {
		return append([]int{via[0]}, merge(via[1:], vib)...)
	}

	return append([]int{vib[0]}, merge(via, vib[1:])...)
}
