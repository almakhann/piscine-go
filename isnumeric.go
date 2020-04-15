package piscine

func IsNumeric(str string) bool {
	for _, l := range str {
		if !(48 <= l && l <= 57) {
			return false
		}
	}
	return true
}
