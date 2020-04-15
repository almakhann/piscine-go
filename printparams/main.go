package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for _, letter := range arguments {
		for _, letter := range letter {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}
