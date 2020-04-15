package piscine

func MakeRange(min, max int) []int {

	var result []int

	if min >= max {
		return result
	}
	s := max - min

	result = make([]int, s)
	for i := 0; i < s; i++ {
		result[i] = min
		min++
	}
	return result
}
