package piscine

func Map(f func(int) bool, arr []int) []bool {
	count := 0
	for range arr {
		count++
	}
	l := make([]bool, count)

	for i, v := range arr {
		l[i] = f(v)
	}
	return l
}
