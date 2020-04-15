package piscine

func UltimateDivMod(a *int, b *int) {
	var c, d int = *a, *b

	*a = c / d
	*b = c % d
}
