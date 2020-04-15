package piscine

func Rot14(str string) string {
	var a string
	for _, l := range str {
		q := l + 14
		if l >= 'a' && l <= 'z' {
			if q > 122 {
				q = 'a' + (q - 123)
			}
			a = a + string(q)

		} else if l >= 'A' && l <= 'Z' {
			if q > 90 {
				q = 'A' + (q - 91)
			}
			a = a + string(q)
		} else {
			a = a + string(l)
		}

	}
	return a

}
