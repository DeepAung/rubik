package rubik

type IRubik interface {
	Rotates(notations []string) error
	Rotate(notation string) error
}

/*
  4
0 1 2 3
  5
*/

var opposite = [6]uint8{2, 3, 0, 1, 5, 4}

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
		ajacentSide{
			SideIndex: 4,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 1,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 5,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 3,
			Positions: Right,
		},
	},
	{
		ajacentSide{
			SideIndex: 4,
			Positions: Bottom,
		},
		ajacentSide{
			SideIndex: 2,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 5,
			Positions: Top,
		},
		ajacentSide{
			SideIndex: 0,
			Positions: Right,
		},
	},
	{
		ajacentSide{
			SideIndex: 4,
			Positions: Right,
		},
		ajacentSide{
			SideIndex: 3,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 5,
			Positions: Right,
		},
		ajacentSide{
			SideIndex: 1,
			Positions: Right,
		},
	},
	{
		ajacentSide{
			SideIndex: 4,
			Positions: Top,
		},
		ajacentSide{
			SideIndex: 0,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 5,
			Positions: Bottom,
		},
		ajacentSide{
			SideIndex: 2,
			Positions: Right,
		},
	},
	{
		ajacentSide{
			SideIndex: 5,
			Positions: Bottom,
		},
		ajacentSide{
			SideIndex: 2,
			Positions: Top,
		},
		ajacentSide{
			SideIndex: 1,
			Positions: Top,
		},
		ajacentSide{
			SideIndex: 0,
			Positions: Top,
		},
	},
	{
		ajacentSide{
			SideIndex: 4,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 1,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 5,
			Positions: Left,
		},
		ajacentSide{
			SideIndex: 3,
			Positions: Right,
		},
	},
}

type rubik struct {
	State [6][3][3]uint8
}

type ajacentSide struct {
	SideIndex uint8
	Positions threeSide
}

type threeSide [3][2]uint8

var (
	Top    threeSide = [3][2]uint8{{0, 0}, {0, 1}, {0, 2}}
	Bottom threeSide = [3][2]uint8{{2, 0}, {2, 1}, {2, 2}}
	Left   threeSide = [3][2]uint8{{0, 0}, {1, 0}, {2, 0}}
	Right  threeSide = [3][2]uint8{{0, 2}, {1, 2}, {2, 2}}
)

type color int

const ( // TODO: change color order to match the above
	White color = iota
	Yellow
	Red
	Blue
	Green
	Orange
)

func InitRubik() IRubik {
	return &rubik{
		State: initialState,
	}
}

func (r *rubik) Rotates(notations []string) error {
	return nil
}

func (r *rubik) Rotate(notation string) error {
	return nil
}
