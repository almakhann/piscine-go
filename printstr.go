package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	list := []byte(str)

	for _, i := range list {

		z01.PrintRune(rune(i))

	}
}
