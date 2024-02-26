package rubik

/*
Front is [1]

  4
0 1 2 3
  5

  r
b w g y
  o
*/

var opposite = [6]uint8{2, 3, 0, 1, 5, 4}

var (
	top    threeSide = [3][2]uint8{{0, 0}, {0, 1}, {0, 2}}
	bottom threeSide = [3][2]uint8{{2, 0}, {2, 1}, {2, 2}}
	left   threeSide = [3][2]uint8{{0, 0}, {1, 0}, {2, 0}}
	right  threeSide = [3][2]uint8{{0, 2}, {1, 2}, {2, 2}}
)

const (
	blue color = iota
	white
	green
	yellow
	red
	orange
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

var notationCharList = map[byte]struct{}{
	F: {}, R: {}, U: {}, L: {}, B: {}, D: {}, M: {}, E: {}, S: {}, X: {}, Y: {}, Z: {},
}

var notationCharToFaceIndex = map[byte]int{
	F: 1,
	R: 2,
	U: 4,
	L: 0,
	B: 3,
	D: 5,
}

// in 6 side has 4 clockwise ajacent 3 edge each edge has position([2]uint8)
var initialState = [6][3][3]uint8{
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

var around = [6][4]ajacentSide{
	{
		ajacentSide{sideIndex: 4, positions: left},
		ajacentSide{sideIndex: 1, positions: left},
		ajacentSide{sideIndex: 5, positions: left},
		ajacentSide{sideIndex: 3, positions: right},
	},
	{
		ajacentSide{sideIndex: 4, positions: bottom},
		ajacentSide{sideIndex: 2, positions: left},
		ajacentSide{sideIndex: 5, positions: top},
		ajacentSide{sideIndex: 0, positions: right},
	},
	{
		ajacentSide{sideIndex: 4, positions: right},
		ajacentSide{sideIndex: 3, positions: left},
		ajacentSide{sideIndex: 5, positions: right},
		ajacentSide{sideIndex: 1, positions: right},
	},
	{
		ajacentSide{sideIndex: 4, positions: top},
		ajacentSide{sideIndex: 0, positions: left},
		ajacentSide{sideIndex: 5, positions: bottom},
		ajacentSide{sideIndex: 2, positions: right},
	},
	{
		ajacentSide{sideIndex: 5, positions: bottom},
		ajacentSide{sideIndex: 2, positions: top},
		ajacentSide{sideIndex: 1, positions: top},
		ajacentSide{sideIndex: 0, positions: top},
	},
	{
		ajacentSide{sideIndex: 4, positions: left},
		ajacentSide{sideIndex: 1, positions: left},
		ajacentSide{sideIndex: 5, positions: left},
		ajacentSide{sideIndex: 3, positions: right},
	},
}
