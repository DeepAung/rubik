package rubik

import (
	"fmt"
	"strings"

	"github.com/DeepAung/rubik/rubik/constant"
	"github.com/DeepAung/rubik/rubik/historymanager"
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
	Sprint() string
	IsSolved() bool
	State() *[6][3][3]uint8
	CycleNumber(notationsStr string) (times int, moves int, err error)

	SetState(state *[6][3][3]uint8, saveHistory bool)
	Rotates(notationsStr string, saveHistory bool) (moves int, err error)
	Rotate(notation *types.Notation, saveHistory bool) error
	RotateInverse(notation *types.Notation, saveHistory bool) error
	Reset(saveHistory bool)

	Undo(times int)
	Redo(times int)

	// for test
	getNotation(notationStr string) (*types.Notation, error)
}

type rubik struct {
	state          [6][3][3]uint8
	historyManager historymanager.IHistoryManager
}

func NewRubik() IRubik {
	return &rubik{
		state:          constant.InitialState,
		historyManager: historymanager.New(),
	}
}

func (r *rubik) Print() {
	fmt.Print(r.Sprint())
}

func (r *rubik) Sprint() string {
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
	return fmt.Sprintf(
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

func (r *rubik) IsSolved() bool {
	return utils.SameState(&r.state, &constant.InitialState)
}

func (r *rubik) State() *[6][3][3]uint8 {
	return &r.state
}

// number of times and moves of rotations that makes rubik reverts back to its original state
// e.g.
// CycleNumber("F")   => 4, 4, nil
// CycleNumber("F F") => 2, 4, nil
// CycleNumber("2F")  => 2, 4, nil
func (r *rubik) CycleNumber(notationsStr string) (times int, moves int, err error) {
	startState := r.state
	LIMIT := 1000

	for {
		rotatedMoves, err := r.Rotates(notationsStr, false)
		if err != nil {
			return 0, 0, fmt.Errorf("Rotates error: %v", err)
		}

		times++
		moves += rotatedMoves

		if utils.SameState(&startState, &r.state) || times >= LIMIT {
			break
		}
	}

	if times == LIMIT {
		return 0, 0, fmt.Errorf("cycle more than %d...", LIMIT)
	}

	return times, moves, nil
}

func (r *rubik) SetState(state *[6][3][3]uint8, saveHistory bool) {
	if saveHistory {
		r.historyManager.UpdateSet(&r.state, state)
	}

	r.state = *state
}

func (r *rubik) Rotates(notationsStr string, saveHistory bool) (moves int, err error) {
	slice := strings.Split(notationsStr, " ")
	for _, notationStr := range slice {
		notation, err := r.getNotation(notationStr)
		if err != nil {
			return 0, err
		}

		moves += int(notation.Number)

		err = r.Rotate(notation, saveHistory)
		if err != nil {
			return 0, err
		}
	}

	return moves, err
}

func (r *rubik) Rotate(notation *types.Notation, saveHistory bool) error {
	faceIndex, ok := constant.NotationCharToFaceIndex[notation.NotationChar]
	if !ok {
		// TODO:
		return fmt.Errorf("notation char is MESXYZ, no implementation of it yet")
	}

	length := int(notation.Number % 4)
	for i := 0; i < length; i++ {
		r.rotateFace(faceIndex, notation.Inverse)
		r.rotateSide(faceIndex, notation.Inverse)
	}

	if saveHistory {
		r.historyManager.UpdateRotate(notation)
	}

	return nil
}

func (r *rubik) RotateInverse(notation *types.Notation, saveHistory bool) error {
	return r.Rotate(&types.Notation{
		Number:       notation.Number,
		NotationChar: notation.NotationChar,
		Inverse:      !notation.Inverse,
	}, saveHistory)
}

func (r *rubik) Reset(saveHistory bool) {
	if saveHistory {
		r.historyManager.UpdateSet(&r.state, &constant.InitialState)
	}

	r.state = constant.InitialState
}

func (r *rubik) Undo(times int) {
	r.historyManager.Undo(times, r)
}

func (r *rubik) Redo(times int) {
	r.historyManager.Redo(times, r)
}

// --------------------------------------------------------------- //

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
