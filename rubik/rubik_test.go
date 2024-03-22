package rubik

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/DeepAung/rubik/rubik/constant"
	"github.com/DeepAung/rubik/rubik/types"
	"github.com/DeepAung/rubik/rubik/utils"
)

func BenchmarkRotates(b *testing.B) {
	rubik := NewRubik()
	input := genNotationInput(10000)

	for i := 0; i < b.N; i++ {
		rubik.Rotates(input, true)
	}
}

func genNotationInput(n int) string {
	notationSet := [12]byte{'F', 'R', 'U', 'L', 'B', 'D', 'M', 'E', 'S', 'X', 'Y', 'Z'}

	var res strings.Builder
	for i := 0; i < n; i++ {
		number := rand.Intn(100)
		char := notationSet[rand.Intn(12)]
		inverse := rand.Intn(2) == 0

		if number != 0 {
			res.WriteString(strconv.Itoa(number))
		}
		res.WriteByte(char)
		if inverse {
			res.WriteByte('\'')
		}

		if i != n-1 {
			res.WriteByte(' ')
		}
	}

	return res.String()
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

const rotateInput = "F R U L B D M E S X Y Z F' R' U' L' B' D' M' E' S' X' Y' Z'"
const B = constant.Blue
const W = constant.White
const G = constant.Green
const Y = constant.Yellow
const R = constant.Red
const O = constant.Orange

var rotateOutput = [6][3][3]uint8{
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

func TestRotates(t *testing.T) {
	rubik := NewRubik()
	_, err := rubik.Rotates(rotateInput, true)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.SameState(rubik.State(), &rotateOutput) {
		t.Fatal("not the same state")
	}
}
