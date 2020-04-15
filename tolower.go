package piscine

func ToLower(s string) string {
	var q = s
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && 'Z' >= s[i] {
			q = q[:i] + string(s[i]+32) + q[i+1:]
		}
	}
	return q
}
