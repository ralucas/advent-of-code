package bit

var HexToBin = map[string][]int8{
	"0": {0, 0, 0, 0},
	"1": {0, 0, 0, 1},
	"2": {0, 0, 1, 0},
	"3": {0, 0, 1, 1},
	"4": {0, 1, 0, 0},
	"5": {0, 1, 0, 1},
	"6": {0, 1, 1, 0},
	"7": {0, 1, 1, 1},
	"8": {1, 0, 0, 0},
	"9": {1, 0, 0, 1},
	"A": {1, 0, 1, 0},
	"B": {1, 0, 1, 1},
	"C": {1, 1, 0, 0},
	"D": {1, 1, 0, 1},
	"E": {1, 1, 1, 0},
	"F": {1, 1, 1, 1},
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
