package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	count := 0

	for range tab {
		count++
	}

	for i := 1; i < count; i++ {
		if f(tab[i], tab[i-1]) < 0 && f(tab[1], tab[0]) > 0 || f(tab[i], tab[i-1]) > 0 && f(tab[1], tab[0]) <= 0 {
			return false
		}
	}
	return true
}
