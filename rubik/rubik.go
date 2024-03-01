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
	Print()
	State() *[6][3][3]uint8
	Rotates(notations ...string) error
	Reset()

	// for test
	getNotation(notationStr string) (*types.Notation, error)
}

type rubik struct {
	state [6][3][3]uint8
}

func NewRubik() IRubik {
	return &rubik{
		state: constant.InitialState,
	}
}

func (r *rubik) Print() {
	str := "\n" +
		"        %s%s%s\n" +
		"        %s%s%s\n" +
		"        %s%s%s\n" +
		"\n" +
		"%s%s%s  %s%s%s  %s%s%s  %s%s%s \n" +
		"%s%s%s  %s%s%s  %s%s%s  %s%s%s \n" +
		"%s%s%s  %s%s%s  %s%s%s  %s%s%s \n" +
		"\n" +
		"        %s%s%s\n" +
		"        %s%s%s\n" +
		"        %s%s%s\n"

	a := &r.state
	b := &constant.IntToColor
	fmt.Printf(
		str,
		b[a[4][0][0]].Sprint("  "),
		b[a[4][0][1]].Sprint("  "),
		b[a[4][0][2]].Sprint("  "),
		b[a[4][1][0]].Sprint("  "),
		b[a[4][1][1]].Sprint("  "),
		b[a[4][1][2]].Sprint("  "),
		b[a[4][2][0]].Sprint("  "),
		b[a[4][2][1]].Sprint("  "),
		b[a[4][2][2]].Sprint("  "),

		b[a[0][0][0]].Sprint("  "),
		b[a[0][0][1]].Sprint("  "),
		b[a[0][0][2]].Sprint("  "),
		b[a[1][0][0]].Sprint("  "),
		b[a[1][0][1]].Sprint("  "),
		b[a[1][0][2]].Sprint("  "),
		b[a[2][0][0]].Sprint("  "),
		b[a[2][0][1]].Sprint("  "),
		b[a[2][0][2]].Sprint("  "),
		b[a[3][0][0]].Sprint("  "),
		b[a[3][0][1]].Sprint("  "),
		b[a[3][0][2]].Sprint("  "),

		b[a[0][1][0]].Sprint("  "),
		b[a[0][1][1]].Sprint("  "),
		b[a[0][1][2]].Sprint("  "),
		b[a[1][1][0]].Sprint("  "),
		b[a[1][1][1]].Sprint("  "),
		b[a[1][1][2]].Sprint("  "),
		b[a[2][1][0]].Sprint("  "),
		b[a[2][1][1]].Sprint("  "),
		b[a[2][1][2]].Sprint("  "),
		b[a[3][1][0]].Sprint("  "),
		b[a[3][1][1]].Sprint("  "),
		b[a[3][1][2]].Sprint("  "),

		b[a[0][2][0]].Sprint("  "),
		b[a[0][2][1]].Sprint("  "),
		b[a[0][2][2]].Sprint("  "),
		b[a[1][2][0]].Sprint("  "),
		b[a[1][2][1]].Sprint("  "),
		b[a[1][2][2]].Sprint("  "),
		b[a[2][2][0]].Sprint("  "),
		b[a[2][2][1]].Sprint("  "),
		b[a[2][2][2]].Sprint("  "),
		b[a[3][2][0]].Sprint("  "),
		b[a[3][2][1]].Sprint("  "),
		b[a[3][2][2]].Sprint("  "),

		b[a[5][0][0]].Sprint("  "),
		b[a[5][0][1]].Sprint("  "),
		b[a[5][0][2]].Sprint("  "),
		b[a[5][1][0]].Sprint("  "),
		b[a[5][1][1]].Sprint("  "),
		b[a[5][1][2]].Sprint("  "),
		b[a[5][2][0]].Sprint("  "),
		b[a[5][2][1]].Sprint("  "),
		b[a[5][2][2]].Sprint("  "),
	)
}

func (r *rubik) State() *[6][3][3]uint8 {
	return &r.state
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
		// TODO:
		fmt.Errorf("notation char is MESXYZ, no implementation of it yet")
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
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
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

	slice := make([]*uint8, 12)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			idx := adjSide[i].SideIndex
			pos := adjSide[i].Positions[j]
			slice[i*3+j] = &r.state[idx][pos[0]][pos[1]]
		}
	}

	if inverse {
		utils.ShiftBy3(slice, false)
	} else {
		utils.ShiftBy3(slice, true)
	}

}

func (r *rubik) Reset() {
	r.state = constant.InitialState
}
