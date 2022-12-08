package maputil

type Any interface {
	string | int
}

func Values[T Any](m map[T]T) []T {
	var output []T
	for _, v := range m {
		output = append(output, v)
	}

	return output
}
