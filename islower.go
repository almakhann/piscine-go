package piscine

func IsLower(str string) bool {
	for _, l := range str {
		if !('a' <= l && 'z' >= l) {
			return false
		}
	}
	return true

}
