package bit

import "fmt"

func Htob(s string) ([]int8, error) {
	switch s {
	case "0": 
		return []int8{0, 0, 0, 0}, nil
	case "1": 
		return []int8{0, 0, 0, 1}, nil
	case "2": 
		return []int8{0, 0, 1, 0}, nil
	case "3": 
		return []int8{0, 0, 1, 1}, nil
	case "4": 
		return []int8{0, 1, 0, 0}, nil
	case "5": 
		return []int8{0, 1, 0, 1}, nil
	case "6": 
		return []int8{0, 1, 1, 0}, nil
	case "7": 
		return []int8{0, 1, 1, 1}, nil
	case "8": 
		return []int8{1, 0, 0, 0}, nil
	case "9": 
		return []int8{1, 0, 0, 1}, nil
	case "A": 
		return []int8{1, 0, 1, 0}, nil
	case "B": 
		return []int8{1, 0, 1, 1}, nil
	case "C": 
		return []int8{1, 1, 0, 0}, nil
	case "D": 
		return []int8{1, 1, 0, 1}, nil
	case "E": 
		return []int8{1, 1, 1, 0}, nil
	case "F": 
		return []int8{1, 1, 1, 1}, nil
	default:
		return []int8{}, fmt.Errorf("invalid character: [%s]", s)
	}

}

func Itob(val int) []int8 {
	bin := make([]int8, 36)
	pow := 1
	i := 0
	for i < 36 {
		if val>>i == 1 {
			bin[35-i] = 1
			val -= pow

			if val == 0 {
				break
			}
			// reset
			i = 0
			pow = 1
			continue
		}

		pow *= 2
		i++
	}

	return bin
}

func Btoi(b []int8) int {
	
	if len(b) == 0 {
		return 0
	}

	r := len(b) - 1

	output := 0
	rs := 1

	for r >= 0 {
		if b[r] == 1 {
			output += rs
		}
		r--
		rs *= 2
	}

	return output
}
