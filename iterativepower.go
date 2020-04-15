package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power > 0 {
		count := nb
		for i := 1; i < power; i++ {
			count = count * nb

		}
		return count
	}
	return 0
}
