package piscine

func IsAlpha(str string) bool {
	for _, l := range str {
		if !(('a' <= l && 'z' >= l) || ('A' <= l && 'Z' >= l) || (48 <= l && l <= 57)) {
			return false
		}
	}
	return true
}
