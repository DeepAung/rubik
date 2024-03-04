package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DeepAung/rubik/rubik"
)

func main() {
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

			err = rubik.Rotates(line, true)
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
