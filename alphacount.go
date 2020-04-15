package piscine

func AlphaCount(str string) int {
	var a int
	for _, l := range str {
		if ('a' <= l && 'z' >= l) || ('A' <= l && 'Z' >= l) {
			a = a + 1
		}
	}
	return a
}
