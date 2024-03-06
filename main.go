package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DeepAung/rubik/rubik"
	"github.com/DeepAung/rubik/ui"
)

func main() {
	//testUndoRedo()
	// testCycleNumber()

	ui.Start()

	// s := "test"
	// x := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7")).Render("a\naaaaa\n")
	// z := "================="
	//
	// fmt.Println(lipgloss.JoinVertical(lipgloss.Top, s, x, z))
}

func testUndoRedo() {
	var mode int
	var times int
	in := bufio.NewReader(os.Stdin)

	rubik := rubik.NewRubik()

	for {
		rubik.Print()

		fmt.Println("=====================================")
		fmt.Println("1. rotate")
		fmt.Println("2. reset")
		fmt.Println("3. undo")
		fmt.Println("4. redo")
		fmt.Scanln(&mode)

		switch mode {
		case 1:
			fmt.Printf("notations: ")
			line, err := in.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			line = strings.TrimSpace(line)

			_, err = rubik.Rotates(line, true)
			if err != nil {
				log.Fatal(err)
			}

		case 2:
			rubik.Reset(true)

		case 3:
			fmt.Printf("times: ")
			fmt.Scanln(&times)
			rubik.Undo(times)

		case 4:
			fmt.Printf("times: ")
			fmt.Scanln(&times)
			rubik.Redo(times)

		}
	}
}

func testCycleNumber() {
	in := bufio.NewReader(os.Stdin)

	rubik := rubik.NewRubik()

	for {
		fmt.Printf("notations: ")
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)

		times, moves, err := rubik.CycleNumber(line)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Cycle times: ", times, " | moves: ", moves)
	}
}
