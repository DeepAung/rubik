package rubik

import (
	"testing"

	"github.com/DeepAung/rubik/rubik/types"
)

type notationTest struct {
	input  string
	output *types.Notation
}

var notationTestList = []notationTest{
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

func TestGetNotation(t *testing.T) {
	rubik := NewRubik()
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
