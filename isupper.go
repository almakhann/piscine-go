package piscine

func IsUpper(str string) bool {
	for _, l := range str {
		if !('A' <= l && 'Z' >= l) {
			return false
		}
	}
	return true

}
