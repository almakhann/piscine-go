package piscine

func SplitWhiteSpaces(str string) []string {

	slice := []rune(str)
	word_count := 0
	rune_count := -1

	for range slice {
		rune_count++
	}

	sp := 0
	for i, r := range slice {
		if i != 0 && i != rune_count {
			if r == ' ' || r == '	' || r == '\n' {
				if i > sp {
					word_count++
				}
				sp = i + 1
			}

		}
	}

	word_count++

	result := make([]string, word_count)

	sp = 0
	n := 0

	for i, r := range slice {
		if r == ' ' || r == '	' || r == '\n' {
			if i > sp {
				result[n] = string(slice[sp:i])
				n++
			}
			sp = i + 1
		}
		if i == rune_count {
			result[n] = string(slice[sp:])
		}
	}
	return result

}
