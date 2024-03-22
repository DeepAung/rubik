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
	RotatesPipeline(notationsStr string, saveHistory bool) (moves int, err error)
	Rotate(notation *types.Notation, saveHistory bool) error
	RotateInverse(notation *types.Notation, saveHistory bool) error
	Reset(saveHistory bool)

	Undo(times int)
	Redo(times int)
	CanUndo() bool
	CanRedo() bool

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
			return 0, 0, err
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
	if notationsStr == "" {
		return 0, fmt.Errorf("no notations")
	}

	savedState := r.state

	slice := strings.Split(notationsStr, " ")
	for _, notationStr := range slice {
		notation, err := r.getNotation(notationStr)
		if err != nil {
			r.state = savedState
			return 0, err
		}

		moves += int(notation.Number)

		err = r.Rotate(notation, saveHistory)
		if err != nil {
			r.state = savedState
			return 0, err
		}
	}

	return moves, err
}

type pipelineRes struct {
	Notation types.Notation
	Error    error
}

func (r *rubik) RotatesPipeline(notationsStr string, saveHistory bool) (moves int, err error) {
	if notationsStr == "" {
		return 0, fmt.Errorf("no notations")
	}

	savedState := r.state

	ch1 := r.getNotationsPipeline(notationsStr, &moves)
	ch2 := r.rotatesPipeline(ch1, saveHistory)

	for err := range ch2 {
		if err != nil {
			r.state = savedState
			return 0, err
		}
	}

	return moves, nil
}

func (r *rubik) getNotationsPipeline(notationsStr string, moves *int) <-chan pipelineRes {
	out := make(chan pipelineRes)

	go func() {
		defer close(out)

		initialRes := types.Notation{
			Number:       0,
			NotationChar: '0',
			Inverse:      false,
		}

		res := initialRes
		initial := true
		collectingDigit := true
		for _, char := range notationsStr {
			if char == ' ' {
				if initial {
					continue
				}
				if res.Number == 0 {
					res.Number = 1
				}
				*moves += int(res.Number)
				out <- pipelineRes{Notation: res}

				res = initialRes
				initial = true
				collectingDigit = true

				continue
			}

			if collectingDigit {
				if utils.IsDigit(char) {
					res.Number = res.Number*10 + uint(char-'0')
					initial = false
				} else if utils.IsNotationChar(char, constant.NotationCharSet) {
					initial = false
					collectingDigit = false
					res.NotationChar = byte(char)
				} else {
					out <- pipelineRes{Error: fmt.Errorf("invalid notation string")}
					return
				}

			} else {
				if char == '\'' { // char == singlequote
					res.Inverse = true
				} else {
					out <- pipelineRes{Error: fmt.Errorf("invalid notation string")}
				}
			}
		}

		if initial {
			return
		}
		if res.Number == 0 {
			res.Number = 1
		}
		*moves += int(res.Number)
		out <- pipelineRes{Notation: res}
	}()

	return out
}

func (r *rubik) rotatesPipeline(in <-chan pipelineRes, saveHistory bool) <-chan error {
	out := make(chan error)

	go func() {
		defer close(out)

		for res := range in {
			if res.Error != nil {
				out <- res.Error
				break
			}

			err := r.Rotate(&res.Notation, saveHistory)
			if err != nil {
				out <- err
				break
			}

			out <- nil
		}
	}()

	return out
}

func (r *rubik) Rotate(notation *types.Notation, saveHistory bool) error {
	if notation == nil {
		return fmt.Errorf("no notation")
	}

	faceIndex, ok := constant.NotationCharToFaceIndex[notation.NotationChar]
	if !ok {
		return fmt.Errorf("invalid notation char")
	}

	length := int(notation.Number % 4)

	switch notation.NotationChar {

	case 'X', 'Y', 'Z':
		tmp := constant.ForXYZ[notation.NotationChar]
		if notation.Inverse {
			tmp[0].Inverse = !tmp[0].Inverse
			tmp[1].Inverse = !tmp[1].Inverse
			tmp[2].Inverse = !tmp[2].Inverse
		}

		for i := 0; i < length; i++ {
			r.rotateFace(tmp[0].Char, tmp[0].Inverse)
			r.rotateFace(tmp[2].Char, tmp[2].Inverse)

			r.rotateSide(tmp[0].Char, tmp[0].Inverse)
			r.rotateSide(tmp[1].Char, tmp[1].Inverse)
			r.rotateSide(tmp[2].Char, tmp[2].Inverse)
		}

	default:
		for i := 0; i < length; i++ {
			if faceIndex != -1 {
				r.rotateFace(notation.NotationChar, notation.Inverse)
			}
			r.rotateSide(notation.NotationChar, notation.Inverse)
		}

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

func (r *rubik) CanUndo() bool {
	return r.historyManager.CanUndo()
}

func (r *rubik) CanRedo() bool {
	return r.historyManager.CanRedo()
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

func (r *rubik) rotateFace(char byte, inverse bool) {
	faceIndex := constant.NotationCharToFaceIndex[char]
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

func (r *rubik) rotateSide(char byte, inverse bool) {
	sides := constant.NotationCharToSides[char]

	slice := make([]*uint8, 12)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			idx := sides[i].SideIndex
			pos := sides[i].Positions[j]
			slice[i*3+j] = &r.state[idx][pos[0]][pos[1]]
		}
	}

	if inverse {
		utils.ShiftBy3(slice, false)
	} else {
		utils.ShiftBy3(slice, true)
	}

}
