package types

type AjacentSide struct {
	SideIndex uint8
	Positions ThreeSide
}

type ThreeSide [3][2]uint8

type Notation struct {
	Number       uint
	NotationChar byte
	Inverse      bool
}

type OneNotation struct {
	Char    byte
	Inverse bool
}

type History struct {
	HistoryType    HistoryType
	NotationChange Notation
	Previous       [6][3][3]uint8
}

type HistoryType uint8

const (
	None   HistoryType = 0
	Rotate HistoryType = 1
	Reset  HistoryType = 2
)
