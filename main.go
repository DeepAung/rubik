package main

import (
	"log"

	"github.com/DeepAung/rubik/rubik"
)

func main() {
	r := rubik.NewRubik()
	r.Print()

	if err := r.Rotates("2F", "L", "L'", "R", "B"); err != nil {
		log.Fatal("error: ", err)
	}
	r.Print()
}
