package rubik

import (
	"testing"

	"github.com/DeepAung/rubik/rubik/constant"
	"github.com/DeepAung/rubik/rubik/types"
	"github.com/DeepAung/rubik/rubik/utils"
)

func BenchmarkRubikNormalRotates(b *testing.B) {
	rubik := NewRubik()
	input := "F R U L B D M E S X Y Z F' R' U' L' B' D' M' E' S' X' Y' Z'"
	for i := 0; i < b.N; i++ {
		rubik.Rotates(input, true)
	}
}

func TestGetNotation(t *testing.T) {
	rubik := NewRubik()

	var notationTestList = []struct {
		input  string
		output *types.Notation
	}{
		{input: "F", output: &types.Notation{Number: 1, NotationChar: 'F', Inverse: false}},
		{input: "1F", output: &types.Notation{Number: 1, NotationChar: 'F', Inverse: false}},
		{input: "F'", output: &types.Notation{Number: 1, NotationChar: 'F', Inverse: true}},
		{input: "1F", output: &types.Notation{Number: 1, NotationChar: 'F', Inverse: false}},
		{input: "1F'", output: &types.Notation{Number: 1, NotationChar: 'F', Inverse: true}},
		{input: "157F", output: &types.Notation{Number: 157, NotationChar: 'F', Inverse: false}},
		{input: "157F'", output: &types.Notation{Number: 157, NotationChar: 'F', Inverse: true}},
		{input: "F127", output: nil},
		{input: "'F127", output: nil},
	}

	for _, item := range notationTestList {
		notation, err := rubik.getNotation(item.input)
		if (err != nil) != (item.output == nil) {
			t.Fatalf("getNotation() error: \nerr = %v\noutput = %v", err, item.output)
		}

		if err != nil {
			return
		}

		if notation.Number != item.output.Number ||
			notation.NotationChar != item.output.NotationChar ||
			notation.Inverse != item.output.Inverse {
			t.Fatalf("getNotation() error: \nget %v\nexpect %v\n", *notation, *item.output)
		}

		if notation != nil {
			t.Logf("notation: %v | error: %v\n", *notation, err)
		} else {
			t.Logf("error: %v\n", err)
		}
	}
}

func TestRotate(t *testing.T) {
	B, W, G, Y, R, O := constant.Blue, constant.White, constant.Green, constant.Yellow, constant.Red, constant.Orange

	rubik := NewRubik()
	input := "F R U L B D M E S X Y Z F' R' U' L' B' D' M' E' S' X' Y' Z'"
	expectState := [6][3][3]uint8{
		{
			{O, O, W},
			{O, G, Y},
			{Y, G, B},
		},
		{
			{B, W, Y},
			{G, Y, R},
			{Y, R, R},
		},
		{
			{G, Y, R},
			{W, B, B},
			{W, Y, O},
		},
		{
			{W, G, W},
			{R, W, B},
			{Y, G, R},
		},
		{
			{G, R, G},
			{W, R, B},
			{O, B, O},
		},
		{
			{R, Y, B},
			{W, O, O},
			{G, O, B},
		},
	}

	// arr := strings.Split(input, " ")
	// for _, item := range arr {
	// 	rubik.Rotates(item, true)
	// 	fmt.Println("notation string: ", item)
	// 	rubik.Print()
	// }

	rubik.Rotates(input, true)
	if !utils.SameState(rubik.State(), &expectState) {
		t.FailNow()
	}
}
