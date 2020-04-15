package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{}

	setPoint(points)

	var a = "x = " + points.x + ", y = " + points.y

	for _, l := range a {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
