package piscine

func StrRev(s string) string {
	a := StrLen(s)
	var new_s string
	for i := (a - 1); i >= 0; i-- {
		new_s = new_s + string(s[i])
	}
	return new_s
}
