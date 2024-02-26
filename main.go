package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s", &input)

	byteString := []byte(input)
	runeString := []rune(input)

	fmt.Println("byteString: ", byteString)
	fmt.Println("runeString: ", runeString)

	fmt.Println("input[0]: ", input[0])
	fmt.Println("[]rune(input)[0]: ", []rune(input)[0])
}
