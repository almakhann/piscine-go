package piscine

func ToUpper(s string) string {
	var q = s
	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && 'z' >= s[i] {
			q = q[:i] + string(s[i]-32) + q[i+1:]
		}
	}
	return q
}
