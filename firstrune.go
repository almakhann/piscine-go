package piscine

func FirstRune(s string) rune {
	var res rune
	for _, l := range s {
		res = rune(l)
		break
	}
	return res
}
