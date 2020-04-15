package piscine

func Swap(a *int, b *int) {
	var c, d int = *a, *b
	*a = d
	*b = c
}
