package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	count := 0
	for range arg {
		count++
	}

	if count == 0 {
		mes, _ := io.Copy(os.Stdout, os.Stdin)

		s := string(mes)

		for _, letter := range s {
			z01.PrintRune(letter)
		}

	}

	for _, txt := range arg {

		file, err := ioutil.ReadFile(txt)

		if err != nil {
			s := "open " + string(txt) + ": no such file or directory"
			for _, letter := range s {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}

		a := string(file)
		for _, letter := range a {
			z01.PrintRune(letter)
		}

	}
}
