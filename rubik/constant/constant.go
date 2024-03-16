package constant

import (
	"github.com/DeepAung/rubik/rubik/types"
	"github.com/DeepAung/rubik/rubik/utils"
	"github.com/gookit/color"
)

/*
Front is [1]

  4
0 1 2 3
  5

  r
b w g y
  o

       top
left  front  right back
      bottom
*/

var Opposite = [6]uint8{2, 3, 0, 1, 5, 4}

// Rev == Reverse
var (
	Row0    types.ThreeSide = [3][2]uint8{{0, 0}, {0, 1}, {0, 2}}
	Row1    types.ThreeSide = [3][2]uint8{{1, 0}, {1, 1}, {1, 2}}
	Row2    types.ThreeSide = [3][2]uint8{{2, 0}, {2, 1}, {2, 2}}
	Col0    types.ThreeSide = [3][2]uint8{{0, 0}, {1, 0}, {2, 0}}
	Col1    types.ThreeSide = [3][2]uint8{{0, 1}, {1, 1}, {2, 1}}
	Col2    types.ThreeSide = [3][2]uint8{{0, 2}, {1, 2}, {2, 2}}
	Row0Rev types.ThreeSide = [3][2]uint8{{0, 2}, {0, 1}, {0, 0}}
	Row1Rev types.ThreeSide = [3][2]uint8{{1, 2}, {1, 1}, {1, 0}}
	Row2Rev types.ThreeSide = [3][2]uint8{{2, 2}, {2, 1}, {2, 0}}
	Col0Rev types.ThreeSide = [3][2]uint8{{2, 0}, {1, 0}, {0, 0}}
	Col1Rev types.ThreeSide = [3][2]uint8{{2, 1}, {1, 1}, {0, 1}}
	Col2Rev types.ThreeSide = [3][2]uint8{{2, 2}, {1, 2}, {0, 2}}
)

// uint8 in the state can be mapped to colors
var IntToColor = [6]color.PrinterFace{
	color.RGB(0, 0, 255, true),     // Blue
	color.RGB(255, 255, 255, true), // White
	color.RGB(0, 255, 0, true),     // Green
	color.RGB(255, 255, 0, true),   // Yellow
	color.RGB(255, 0, 0, true),     // Red
	color.RGB(255, 100, 0, true),   // Orange
}

const (
	Blue uint8 = iota
	White
	Green
	Yellow
	Red
	Orange
)

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

var NotationCharSet = utils.MakeSet[byte](F, R, U, L, B, D, M, E, S, X, Y, Z)

var NotationCharToFaceIndex = map[byte]int{
	F: 1, R: 2, U: 4, L: 0, B: 3, D: 5, M: -1, E: -1, S: -1, X: -1, Y: -1, Z: -1,
}

// in 6 face has 4 clockwise ajacents that have 3 edges. each edge has position([2]uint8)
// Warning: make sure to not change this constant
var NotationCharToSides = map[byte]*[4]types.AjacentSide{
	L: {
		types.AjacentSide{SideIndex: 4, Positions: Col0},
		types.AjacentSide{SideIndex: 1, Positions: Col0},
		types.AjacentSide{SideIndex: 5, Positions: Col0},
		types.AjacentSide{SideIndex: 3, Positions: Col2Rev},
	},
	F: {
		types.AjacentSide{SideIndex: 4, Positions: Row2},
		types.AjacentSide{SideIndex: 2, Positions: Col0},
		types.AjacentSide{SideIndex: 5, Positions: Row0Rev},
		types.AjacentSide{SideIndex: 0, Positions: Col2Rev},
	},
	R: {
		types.AjacentSide{SideIndex: 4, Positions: Col2Rev},
		types.AjacentSide{SideIndex: 3, Positions: Col0},
		types.AjacentSide{SideIndex: 5, Positions: Col2Rev},
		types.AjacentSide{SideIndex: 1, Positions: Col2Rev},
	},
	B: {
		types.AjacentSide{SideIndex: 4, Positions: Row0Rev},
		types.AjacentSide{SideIndex: 0, Positions: Col0},
		types.AjacentSide{SideIndex: 5, Positions: Row2},
		types.AjacentSide{SideIndex: 2, Positions: Col2Rev},
	},
	U: {
		types.AjacentSide{SideIndex: 3, Positions: Row0Rev},
		types.AjacentSide{SideIndex: 2, Positions: Row0Rev},
		types.AjacentSide{SideIndex: 1, Positions: Row0Rev},
		types.AjacentSide{SideIndex: 0, Positions: Row0Rev},
	},
	D: {
		types.AjacentSide{SideIndex: 1, Positions: Row2},
		types.AjacentSide{SideIndex: 2, Positions: Row2},
		types.AjacentSide{SideIndex: 3, Positions: Row2},
		types.AjacentSide{SideIndex: 0, Positions: Row2},
	},
	M: {
		types.AjacentSide{SideIndex: 4, Positions: Col1},
		types.AjacentSide{SideIndex: 1, Positions: Col1},
		types.AjacentSide{SideIndex: 5, Positions: Col1},
		types.AjacentSide{SideIndex: 3, Positions: Col1Rev},
	},
	E: {
		types.AjacentSide{SideIndex: 1, Positions: Row1},
		types.AjacentSide{SideIndex: 2, Positions: Row1},
		types.AjacentSide{SideIndex: 3, Positions: Row1},
		types.AjacentSide{SideIndex: 0, Positions: Row1},
	},
	S: {
		types.AjacentSide{SideIndex: 4, Positions: Row1},
		types.AjacentSide{SideIndex: 2, Positions: Col1},
		types.AjacentSide{SideIndex: 5, Positions: Row1Rev},
		types.AjacentSide{SideIndex: 0, Positions: Col1Rev},
	},
}

var ForXYZ = map[byte][3]types.OneNotation{
	X: {
		{Char: 'R', Inverse: false},
		{Char: 'M', Inverse: true},
		{Char: 'L', Inverse: true},
	},
	Y: {
		{Char: 'U', Inverse: false},
		{Char: 'E', Inverse: true},
		{Char: 'D', Inverse: true},
	},
	Z: {
		{Char: 'F', Inverse: false},
		{Char: 'S', Inverse: false},
		{Char: 'B', Inverse: true},
	},
}
