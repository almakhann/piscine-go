package piscine

func IsPrintable(str string) bool {
	for _, l := range str {
		if !(32 <= l && 126 >= l) {
			return false
		}
	}
	return true
}
