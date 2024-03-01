package utils

func MakeSet[T comparable](items ...T) map[T]struct{} {
	result := map[T]struct{}{}
	for _, item := range items {
		result[item] = struct{}{}
	}

	return result
}

func IsDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func IsNotationChar(char rune, notationCharSet map[byte]struct{}) bool {
	if char < 0 || char > 255 {
		return false
	}

	_, ok := notationCharSet[byte(char)]
	return ok
}

// toRight [0, 1, 2, 3, 4, 5, 6, 7] => [5, 6, 7, 0, 1, 2, 3, 4]
// toLeft  [0, 1, 2, 3, 4, 5, 6, 7] => [3, 4, 5, 6, 7, 1, 2, 3]
func ShiftBy3(slice []*uint8, toRight bool) {
	n := len(slice)
	tmp := make([]uint8, 3)

	if toRight {
		tmp[0] = *slice[n-1]
		tmp[1] = *slice[n-2]
		tmp[2] = *slice[n-3]
		for i := n - 1; i > 2; i-- {
			*slice[i] = *slice[i-3]
		}
		*slice[2] = tmp[0]
		*slice[1] = tmp[1]
		*slice[0] = tmp[2]
	} else {
		tmp[0] = *slice[0]
		tmp[1] = *slice[1]
		tmp[2] = *slice[2]
		for i := 0; i < n-3; i++ {
			*slice[i] = *slice[i+3]
		}
		*slice[n-3] = tmp[0]
		*slice[n-2] = tmp[1]
		*slice[n-1] = tmp[2]
	}
}
