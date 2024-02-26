package rubik

type rubik struct {
	state [6][3][3]uint8
}

type ajacentSide struct {
	sideIndex uint8
	positions threeSide
}

type threeSide [3][2]uint8

type color int

type notation struct {
	number       uint
	notationChar byte
	inverse      bool
}
