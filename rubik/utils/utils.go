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
