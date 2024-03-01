package rubik

import (
	"fmt"

	"github.com/DeepAung/rubik/rubik/constant"
	"github.com/DeepAung/rubik/rubik/types"
	"github.com/DeepAung/rubik/rubik/utils"
)

/*
Front is [1]

  4
0 1 2 3
  5

  r
b w g y
  o
*/

type IRubik interface {
	Rotates(notations ...string) error
	rotate(notationStr string) error
	getNotation(notationStr string) (*types.Notation, error)
	rotateFace(faceIndex int, inverse bool)
	rotateSide(faceIndex int, inverse bool)
	Reset()
}

type rubik struct {
	state [6][3][3]uint8
}

func NewRubik() IRubik {
	return &rubik{
		state: constant.InitialState,
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

	faceIndex, ok := constant.NotationCharToFaceIndex[notation.NotationChar]
	if !ok {
		// notation char is MESXYZ
	}

	length := int(notation.Number % 4)
	for i := 0; i < length; i++ {
		r.rotateFace(faceIndex, notation.Inverse)
		r.rotateSide(faceIndex, notation.Inverse)
	}

	return nil
}

func (r *rubik) getNotation(notationStr string) (*types.Notation, error) {
	res := &types.Notation{
		Number:       0,
		NotationChar: '0',
		Inverse:      false,
	}

	collectingDigit := true
	for _, char := range notationStr {
		if collectingDigit {
			if utils.IsDigit(char) {
				res.Number = res.Number*10 + uint(char-'0')
			} else if utils.IsNotationChar(char, constant.NotationCharSet) {
				collectingDigit = false
				res.NotationChar = byte(char)
			} else {
				return nil, fmt.Errorf("invalid notation string")
			}

		} else {
			if char == '\'' { // char == singlequote
				res.Inverse = true
			} else {
				return nil, fmt.Errorf("invalid notation string")
			}
		}
	}

	if res.Number == 0 {
		res.Number = 1
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
	adjSide := &constant.Around[faceIndex]
	_ = adjSide
}

func (r *rubik) Reset() {
	r.state = constant.InitialState
}
