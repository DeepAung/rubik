package utils_test

import (
	"fmt"
	"testing"

	"github.com/DeepAung/rubik/rubik/utils"
)

func TestShiftBy3(t *testing.T) {
	if !testShiftBy3([]uint8{1, 2, 3, 4, 5, 6, 7}, true, []uint8{5, 6, 7, 1, 2, 3, 4, 5, 6, 7}) {
		t.FailNow()
	}

	if !testShiftBy3([]uint8{1, 2, 3, 4, 5, 6, 7}, false, []uint8{4, 5, 6, 7, 1, 2, 3}) {
		t.FailNow()
	}

}

func testShiftBy3(input []uint8, toRight bool, output []uint8) bool {
	slice := make([]*uint8, len(input))
	for i := 0; i < len(input); i++ {
		slice[i] = &input[i]
	}
	utils.ShiftBy3(slice, toRight)

	fmt.Printf("result: ")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", input[i])
	}
	fmt.Printf("\nexpect: ")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", output[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(slice); i++ {
		if input[i] != output[i] {
			fmt.Println("heyyyyy")
			return false
		}
	}

	return true
}
