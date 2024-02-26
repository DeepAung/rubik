package rubik

import "fmt"

type IRubik interface {
	Rotates(notations ...string) error
	rotate(notationStr string) error
	getNotation(notationStr string) (*notation, error)
	rotateFace(faceIndex int, inverse bool)
	rotateSide(faceIndex int, inverse bool)
	Reset()

	// utils
	isDigit(char rune) bool
	isNotationChar(char rune) bool
}

/*
Front is [1]

  4
0 1 2 3
  5

  r
b w g y
  o
*/

func NewRubik() IRubik {
	return &rubik{
		state: initialState,
	}
}

func (r *rubik) Rotates(notations ...string) error {
	for _, notation := range notations {
		err := r.rotate(notation)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rubik) rotate(notationStr string) error {
	notation, err := r.getNotation(notationStr)
	if err != nil {
		return err
	}

	faceIndex, ok := notationCharToFaceIndex[notation.notationChar]
	if !ok {
		// notation char is MESXYZ
	}

	length := int(notation.number % 4)
	for i := 0; i < length; i++ {
		r.rotateFace(faceIndex, notation.inverse)
		r.rotateSide(faceIndex, notation.inverse)
	}

	return nil
}

func (r *rubik) getNotation(notationStr string) (*notation, error) {
	res := &notation{
		number:       0,
		notationChar: '0',
		inverse:      false,
	}

	collectingDigit := true
	for _, char := range notationStr {
		if collectingDigit {
			if r.isDigit(char) {
				res.number = res.number*10 + uint(char-'0')
			} else if r.isNotationChar(char) {
				collectingDigit = false
				res.notationChar = byte(char)
			} else {
				return nil, fmt.Errorf("invalid notation string")
			}

		} else {
			if char == '\'' { // char == singlequote
				res.inverse = true
			} else {
				return nil, fmt.Errorf("invalid notation string")
			}
		}
	}

	if res.number == 0 {
		res.number = 1
	}

	return res, nil
}

func (r *rubik) rotateFace(faceIndex int, inverse bool) {
	arr := &r.state[faceIndex]

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			arr[i][j] = arr[j][i]
		}
	}

	if inverse {
		for j := 0; j < 3; j++ {
			arr[0][j], arr[2][j] = arr[2][j], arr[0][j]
		}
	} else {
		for i := 0; i < 3; i++ {
			arr[i][0], arr[i][2] = arr[i][2], arr[i][0]
		}
	}
}

func (r *rubik) rotateSide(faceIndex int, inverse bool) {
	// adjSide := &around[faceIndex]
}

func (r *rubik) Reset() {
	r.state = initialState
}

// utils ---------------------------------------------------- //
func (r *rubik) isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func (r *rubik) isNotationChar(char rune) bool {
	if char < 0 || char > 255 {
		return false
	}

	_, ok := notationCharList[byte(char)]
	return ok
}
