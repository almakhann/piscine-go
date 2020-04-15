package piscine

func Power10(p int) int {
	r := 1
	if p == 0 {
		return 1
	} else {
		i := 1
		for i <= p {
			r = r * 10
			i++
		}
		return r
	}
}
