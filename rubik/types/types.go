package types

type AjacentSide struct {
	SideIndex uint8
	Positions ThreeSide
}

type ThreeSide [3][2]uint8

type Color int

type Notation struct {
	Number       uint
	NotationChar byte
	Inverse      bool
}
