package constant

import (
	"github.com/DeepAung/rubik/rubik/types"
	"github.com/DeepAung/rubik/rubik/utils"
	"github.com/gookit/color"
)

// TODO: check if all constant is correct

/*
Front is [1]

  4
0 1 2 3
  5

  r
b w g y
  o

       top
Left  front  right back
      bottom
*/

var Opposite = [6]uint8{2, 3, 0, 1, 5, 4}

// TODO: should have TopPrime, etc. and update the Around
var (
	Top    types.ThreeSide = [3][2]uint8{{0, 0}, {0, 1}, {0, 2}}
	Bottom types.ThreeSide = [3][2]uint8{{2, 0}, {2, 1}, {2, 2}}
	Left   types.ThreeSide = [3][2]uint8{{0, 0}, {1, 0}, {2, 0}}
	Right  types.ThreeSide = [3][2]uint8{{0, 2}, {1, 2}, {2, 2}}
)

const (
	Blue types.Color = iota
	White
	Green
	Yellow
	Red
	Orange
)

var IntToColor = [6]color.PrinterFace{
	color.BgBlue,
	color.BgWhite,
	color.BgGreen,
	color.BgYellow,
	color.BgRed,
	color.RGB(255, 165, 0, true),
}

const (
	F byte = 'F'
	R byte = 'R'
	U byte = 'U'
	L byte = 'L'
	B byte = 'B'
	D byte = 'D'
	M byte = 'M'
	E byte = 'E'
	S byte = 'S'
	X byte = 'X'
	Y byte = 'Y'
	Z byte = 'Z'
)

var NotationCharSet = utils.MakeSet[byte](F, R, U, L, B, D, M, E, S, X, Y, Z)

var NotationCharToFaceIndex = map[byte]int{F: 1, R: 2, U: 4, L: 0, B: 3, D: 5}

// in 6 face has 4 clockwise ajacent 3 edge each edge has position([2]uint8)
var InitialState = [6][3][3]uint8{
	{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	},
	{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	},
	{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	},
	{
		{3, 3, 3},
		{3, 3, 3},
		{3, 3, 3},
	},
	{
		{4, 4, 4},
		{4, 4, 4},
		{4, 4, 4},
	},
	{
		{5, 5, 5},
		{5, 5, 5},
		{5, 5, 5},
	},
}

// around in clockwise
var Around = [6][4]types.AjacentSide{
	{
		types.AjacentSide{SideIndex: 4, Positions: Left},
		types.AjacentSide{SideIndex: 1, Positions: Left},
		types.AjacentSide{SideIndex: 5, Positions: Left},
		types.AjacentSide{SideIndex: 3, Positions: Right},
	},
	{
		types.AjacentSide{SideIndex: 4, Positions: Bottom},
		types.AjacentSide{SideIndex: 2, Positions: Left},
		types.AjacentSide{SideIndex: 5, Positions: Top},
		types.AjacentSide{SideIndex: 0, Positions: Right},
	},
	{
		types.AjacentSide{SideIndex: 4, Positions: Right},
		types.AjacentSide{SideIndex: 3, Positions: Left},
		types.AjacentSide{SideIndex: 5, Positions: Right},
		types.AjacentSide{SideIndex: 1, Positions: Right},
	},
	{
		types.AjacentSide{SideIndex: 4, Positions: Top},
		types.AjacentSide{SideIndex: 0, Positions: Left},
		types.AjacentSide{SideIndex: 5, Positions: Bottom},
		types.AjacentSide{SideIndex: 2, Positions: Right},
	},
	{
		types.AjacentSide{SideIndex: 5, Positions: Bottom},
		types.AjacentSide{SideIndex: 2, Positions: Top},
		types.AjacentSide{SideIndex: 1, Positions: Top},
		types.AjacentSide{SideIndex: 0, Positions: Top},
	},
	{
		types.AjacentSide{SideIndex: 4, Positions: Left},
		types.AjacentSide{SideIndex: 1, Positions: Left},
		types.AjacentSide{SideIndex: 5, Positions: Left},
		types.AjacentSide{SideIndex: 3, Positions: Right},
	},
}
