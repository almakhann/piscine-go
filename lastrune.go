package piscine

func LastRune(s string) rune {
	var res rune
	len_s := (StrLen(s) - 1)

	res = rune(s[len_s])

	return res
}
