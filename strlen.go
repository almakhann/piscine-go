package piscine

func StrLen(str string) int {
	var count int
	word := []rune(str)
	for range word {
		count = count + 1
	}
	return count
}
