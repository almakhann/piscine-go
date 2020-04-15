package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	i, result := 1, 1
	res := 0
	for result < nb {
		i++
		result = i * i
		if result == nb {
			res = i
		}
	}
	return res
}
