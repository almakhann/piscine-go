package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	} else {
		count := 0
		for start != 1 {
			count++
			if start%2 == 1 {
				start = start*3 + 1
			} else {
				start = start / 2
			}
		}
		return count
	}
}
