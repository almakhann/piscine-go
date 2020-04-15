package piscine

func Index(s string, toFind string) int {
	len_find := StrLen(toFind)
	len_s := StrLen(s)
	for index := 0; index < (len_s - len_find); index++ {
		if s[index:index+len_find] == toFind {
			return index
		}
	}
	return -1
}
