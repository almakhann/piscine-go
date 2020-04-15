package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	count := 0
	for range arguments {
		count = count + 1
	}

	for i := (count - 1); i >= 0; i-- {
		for _, letter := range arguments[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}

}
