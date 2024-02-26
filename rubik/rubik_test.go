package rubik

import (
	"testing"
)

type notationTest struct {
	input  string
	output *notation
}

var notationTestList = []notationTest{
	{input: "F", output: &notation{number: 1, notationChar: 'F', inverse: false}},
	{input: "1F", output: &notation{number: 1, notationChar: 'F', inverse: false}},
	{input: "F'", output: &notation{number: 1, notationChar: 'F', inverse: true}},
	{input: "1F", output: &notation{number: 1, notationChar: 'F', inverse: false}},
	{input: "1F'", output: &notation{number: 1, notationChar: 'F', inverse: true}},
	{input: "157F", output: &notation{number: 157, notationChar: 'F', inverse: false}},
	{input: "157F'", output: &notation{number: 157, notationChar: 'F', inverse: true}},
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

		if notation.number != item.output.number ||
			notation.notationChar != item.output.notationChar ||
			notation.inverse != item.output.inverse {
			t.Fatalf("getNotation() error: \nget %v\nexpect %v\n", *notation, *item.output)
		}

		if notation != nil {
			t.Logf("notation: %v | error: %v\n", *notation, err)
		} else {
			t.Logf("error: %v\n", err)
		}
	}
}
