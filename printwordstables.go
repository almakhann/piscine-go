package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {

	for _, w := range table {
		for _, r := range w {

			z01.PrintRune(r)

		}
		z01.PrintRune('\n')
	}

}
