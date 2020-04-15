package piscine

func Abort(a, b, c, d, e int) int {
	l := []int{a, b, c, d, e}
	count := 0

	for range l {
		count++
	}

	q := count
	for q >= 0 {
		for i := 0; i < count-1; i++ {
			if l[i] > l[i+1] {
				l[i], l[i+1] = l[i+1], l[i]
			}
		}
		q--
	}

	return l[count/2]
}
